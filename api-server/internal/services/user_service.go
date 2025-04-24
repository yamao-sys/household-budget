package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/internal/validators"
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	ValidateSignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error
	SignUp(ctx context.Context, requestParams api.PostUsersSignUpJSONRequestBody) error
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

// NOTE: パスワードの文字列をハッシュ化する
func (us *userService) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
