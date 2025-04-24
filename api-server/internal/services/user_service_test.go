package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/test/factories"
	"fmt"
	"net/http"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestUserServiceSuite struct {
	WithDBSuite
}

var testUserService UserService

func (s *TestUserServiceSuite) SetupTest() {
	s.SetDBCon()

	testUserService = NewUserService(DBCon)
}

func (s *TestUserServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestUserServiceSuite) TestValidateSignUp_SuccessRequiredFields() {
	requestParams := api.PostUsersSignUpJSONRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testUserService.ValidateSignUp(ctx, requestParams)

	assert.Nil(s.T(), result)
}

func (s *TestUserServiceSuite) TestValidateSignUp_ValidationErrorRequiredFields() {
	requestParams := api.PostUsersSignUpJSONRequestBody{
		Name: "",
		Email: "",
		Password: "",
	}

	result := testUserService.ValidateSignUp(ctx, requestParams)

	assert.NotNil(s.T(), result)
	if errors, ok := result.(validation.Errors); ok {
		for field, err := range errors {
			message := err.Error()
			switch field {
			case "name":
				assert.Equal(s.T(), "ユーザ名は必須入力です。", message)
			case "email":
				assert.Equal(s.T(), "Emailは必須入力です。", message)
			case "password":
				assert.Equal(s.T(), "パスワードは必須入力です。", message)
			}
		}
	}
}

func (s *TestUserServiceSuite) TestSignUp_SuccessRequiredFields() {
	requestParams := api.PostUsersSignUpJSONRequestBody{
		Name: "name",
		Email: "test@example.com",
		Password: "Password",
	}

	result := testUserService.SignUp(ctx, requestParams)

	assert.Nil(s.T(), result)

	// NOTE: ユーザが作成されていることを確認
	var exists bool
	err := DBCon.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", "test@example.com").Find(&exists).Error
	if err != nil {
		fmt.Println("err", err)
	}
	assert.True(s.T(), exists)
}

func (s *TestUserServiceSuite) TestSignIn_StatusOK() {
	// NOTE: テスト用ユーザの作成
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostUsersSignInJSONRequestBody{Email: "test@example.com", Password: "password"}

	statusCode, tokenString, err := testUserService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int(http.StatusOK), statusCode)
	assert.NotNil(s.T(), tokenString)
	assert.Nil(s.T(), err)
}

func (s *TestUserServiceSuite) TestSignIn_BadRequest() {
	// NOTE: テスト用ユーザの作成
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostUsersSignInJSONRequestBody{Email: "test_@example.com", Password: "password"}

	statusCode, tokenString, err := testUserService.SignIn(ctx, requestParams)

	assert.Equal(s.T(), int(http.StatusBadRequest), statusCode)
	assert.Equal(s.T(), "", tokenString)
	assert.Equal(s.T(), "メールアドレスまたはパスワードに該当するユーザが存在しません。", err.Error())
}

func TestUserService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestUserServiceSuite))
}
