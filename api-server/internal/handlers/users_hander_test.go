package handlers

import (
	api "apps/api"
	"apps/internal/models"
	"apps/test/factories"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/oapi-codegen/testutil"
)

type TestUsersHandlerSuite struct {
	WithDBSuite
}

func (s *TestUsersHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers()

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestUsersHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestUsersHandlerSuite) TestPostUserValidateSignUp_SuccessRequiredFields() {
	reqBody := api.UserSignUpInput{
		Name: "test_name",
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/validateSignUp").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.UserSignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))
}

func (s *TestUsersHandlerSuite) TestPostUserValidateSignUp_ValidationErrorRequiredFields() {
	reqBody := api.UserSignUpInput{
		Name: "",
		Email: "",
		Password: "",
	}
	result := testutil.NewRequest().Post("/users/validateSignUp").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.UserSignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"ユーザ名は必須入力です。"}, res.Errors.Name)
	assert.Equal(s.T(), &[]string{"Emailは必須入力です。"}, res.Errors.Email)
	assert.Equal(s.T(), &[]string{"パスワードは必須入力です。"}, res.Errors.Password)
}

func (s *TestUsersHandlerSuite) TestPostUserValidateSignUp_StatusForbidden() {
	reqBody := api.UserSignUpInput{
		Name: "test_name",
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/validateSignUp").WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: ユーザが作成されていないことを確認
	var exists bool
	DBCon.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", "test@example.com").Find(&exists)
	assert.False(s.T(), exists)
}

func (s *TestUsersHandlerSuite) TestPostUserSignUp_SuccessRequiredFields() {
	reqBody := api.UserSignUpInput{
		Name: "test_name",
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signUp").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.UserSignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	jsonErrors, _ := json.Marshal(res.Errors)
	assert.Equal(s.T(), "{}", string(jsonErrors))

	// NOTE: ユーザが作成されていることを確認
	var exists bool
	DBCon.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", "test@example.com").Find(&exists)
	if err != nil {
		fmt.Println("err", err)
	}
	assert.True(s.T(), exists)
}

func (s *TestUsersHandlerSuite) TestPostUserSignUp_ValidationErrorRequiredFields() {
	reqBody := api.UserSignUpInput{
		Name: "",
		Email: "",
		Password: "",
	}
	result := testutil.NewRequest().Post("/users/signUp").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.UserSignUpResponseJSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), int(http.StatusOK), res.Code)
	assert.Equal(s.T(), &[]string{"ユーザ名は必須入力です。"}, res.Errors.Name)
	assert.Equal(s.T(), &[]string{"Emailは必須入力です。"}, res.Errors.Email)
	assert.Equal(s.T(), &[]string{"パスワードは必須入力です。"}, res.Errors.Password)

	// NOTE: ユーザが作成されていないことを確認
	var exists bool
	DBCon.Model(&models.User{}).Select("count(*) > 0").Find(&exists)
	if err != nil {
		fmt.Println("err", err)
	}
	assert.False(s.T(), exists)
}

func (s *TestUsersHandlerSuite) TestPostUserSignUp_StatusForbidden() {
	reqBody := api.UserSignUpInput{
		Name: "test_name",
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signUp").WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	// NOTE: ユーザが作成されていないことを確認
	var exists bool
	DBCon.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", "test@example.com").Find(&exists)
	assert.False(s.T(), exists)
}

func (s *TestUsersHandlerSuite) TestPostUsersSignIn_StatusOk() {
	// NOTE: テスト用ユーザの作成
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	reqBody := api.UserSignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	cookieString := result.Recorder.Result().Header.Values("Set-Cookie")[0]
	assert.NotEmpty(s.T(), cookieString)
}

func (s *TestUsersHandlerSuite) TestPostUserSignIn_BadRequest() {
	// NOTE: テスト用ユーザの作成
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	reqBody := api.UserSignInInput{
		Email: "test_@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), int(http.StatusBadRequest), result.Code())

	var res api.UserSignInBadRequestResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), []string{"メールアドレスまたはパスワードに該当するユーザが存在しません。"}, res.Errors)
}

func (s *TestUsersHandlerSuite) TestPostUserSignIn_StatusForbidden() {
	// NOTE: テスト用ユーザの作成
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	reqBody := api.UserSignInInput{
		Email: "test_@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signIn").WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)

	assert.Equal(s.T(), http.StatusForbidden, result.Code())
}

func (s *TestUsersHandlerSuite) TestGetUsersCheckSignedIn_isSignedIn_StatusOk() {
	_, cookieString := s.signIn()

	result := testutil.NewRequest().Get("/users/checkSignedIn").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetUsersCheckSignedIn200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")
	
	assert.True(s.T(), res.IsSignedIn)
}

func (s *TestUsersHandlerSuite) TestGetUsersCheckSignedIn_isNotSignedIn_StatusOk() {
	result := testutil.NewRequest().Get("/users/checkSignedIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func TestUsersHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestUsersHandlerSuite))
}
