package handlers

import (
	api "apps/api"
	"apps/internal/models"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
)

type TestIncomesHandlerSuite struct {
	WithDBSuite
}

func (s *TestIncomesHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers()

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestIncomesHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestIncomesHandlerSuite) TestPostIncomes_Success_StatusOk() {
	user, cookieString := s.signIn()

	reqBody := api.StoreIncomeInput{
		ReceivedAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		ClientName: "client name",
	}

	result := testutil.NewRequest().Post("/incomes").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.PostIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), reqBody.ReceivedAt.Format("2006-01-02"), res.Income.ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), reqBody.Amount, res.Income.Amount)
	assert.Equal(s.T(), reqBody.ClientName, res.Income.ClientName)

	// NOTE: バリデーションエラーがないことを確認
	assert.Equal(s.T(), api.StoreIncomeValidationError{}, res.Errors)

	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.True(s.T(), exists)
}

func (s *TestIncomesHandlerSuite) TestPostIncomes_ValidationErrorRequiredFields_StatusOk() {
	user, cookieString := s.signIn()

	reqBody := api.StoreIncomeInput{}

	result := testutil.NewRequest().Post("/incomes").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.PostIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	amountErrorMessages := []string{"金額は必須入力です。"}
	clientNameErrorMessages := []string{"顧客名は必須入力です。"}
	assert.ElementsMatch(s.T(), amountErrorMessages, *res.Errors.Amount)
	assert.ElementsMatch(s.T(), clientNameErrorMessages, *res.Errors.ClientName)

	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func (s *TestIncomesHandlerSuite) TestPostIncomes_ValidationErrorRequiredFields_StatusUnauthorized() {
	user, _ := s.signIn()

	reqBody := api.StoreIncomeInput{
		ReceivedAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		ClientName: "client name",
	}

	result := testutil.NewRequest().Post("/incomes").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())

	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func (s *TestIncomesHandlerSuite) TestPostIncomes_Success_StatusForbidden() {
	user, cookieString := s.signIn()

	reqBody := api.StoreIncomeInput{
		ReceivedAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		ClientName: "client name",
	}

	result := testutil.NewRequest().Post("/incomes").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusForbidden, result.Code())

	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func TestIncomesHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestIncomesHandlerSuite))
}
