package services

import (
	api "apps/apis"
	"apps/internal/models"
	"apps/internal/validators"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	ValidateSignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error
	SignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error
	SignIn(ctx context.Context, requestParams api.PostUsersSignInJSONRequestBody) (statusCode int, tokenString string, error error)
	ExistsUser(id int) bool
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db}
}

func (us *userService) ValidateSignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error {
	return validators.ValidateSignUp(&requestParams)
}

func (us *userService) SignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error {
	user := models.User{}
	user.Name = requestParams.Name
	user.Email = requestParams.Email

	// NOTE: パスワードをハッシュ化の上、Create処理
	hashedPassword, err := us.encryptPassword(requestParams.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	if err := us.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (us *userService) SignIn(ctx context.Context, requestParams api.PostUsersSignInJSONRequestBody) (statusCode int, tokenString string, error error) {
	// NOTE: emailからの取得
	var user models.User
	if err := us.db.Where("email = ?", requestParams.Email).First(&user).Error; err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "ユーザ")
	}

	// NOTE: パスワードの照合
	if err := us.compareHashPassword(user.Password, requestParams.Password); err != nil {
		return http.StatusBadRequest, "", fmt.Errorf("メールアドレスまたはパスワードに該当する%sが存在しません。", "ユーザ")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_KEY")))
	if err != nil {
		return http.StatusInternalServerError, "", err
	}
	return http.StatusOK, tokenString, nil
}

func (us *userService) ExistsUser(id int) bool {
	var exists bool
	us.db.Model(&models.User{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists)
	return exists
}

// NOTE: パスワードの照合
func (us *userService) compareHashPassword(hashedPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(requestPassword)); err != nil {
		return err
	}
	return nil
}

// NOTE: パスワードの文字列をハッシュ化する
func (us *userService) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
