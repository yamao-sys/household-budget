package handlers

import (
	api "apps/api"
	"apps/internal/models"
	"apps/test/factories"
	"net/http"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/oapi-codegen/testutil"
)

type TestExpensesHandlerSuite struct {
	WithDBSuite
}

func (s *TestExpensesHandlerSuite) SetupTest() {
	s.SetDBCon()

	s.initializeHandlers()

	// NOTE: CSRFトークンのセット
	s.SetCsrfHeaderValues()
}

func (s *TestExpensesHandlerSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithoutBeginngOfMonth_StatusOk() {
	user, cookieString := s.signIn()

	expense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Expense)
	DBCon.Create(&expense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Expense)
	DBCon.Create(&otherExpense)

	result := testutil.NewRequest().Get("/expenses").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 1, len(*res.Expenses))
	assert.Equal(s.T(), expense.PaidAt.Format("2006-01-02"), (*res.Expenses)[0].Date.Format("2006-01-02"))
	assert.Equal(s.T(), "expense", (*res.Expenses)[0].ExtendProps.Type)
	assert.Equal(s.T(), expense.Amount, (*res.Expenses)[0].ExtendProps.Amount)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithBeginngOfMonth_StatusOk() {
	user, cookieString := s.signIn()

	beginningOfMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&beginningOfMonthExpense)
	endOfMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 30, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&endOfMonthExpense)
	endOfPreviousMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&endOfPreviousMonthExpense)
	beginningOfNextMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 5, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&beginningOfNextMonthExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherBeginningOfMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherBeginningOfMonthExpense)

	beginningOfMonth := "2025-04-01"
	result := testutil.NewRequest().Get("/expenses?beginningOfMonth="+beginningOfMonth).WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(*res.Expenses))
	assert.Equal(s.T(), beginningOfMonthExpense.PaidAt.Format("2006-01-02"), (*res.Expenses)[0].Date.Format("2006-01-02"))
	assert.Equal(s.T(), "expense", (*res.Expenses)[0].ExtendProps.Type)
	assert.Equal(s.T(), beginningOfMonthExpense.Amount, (*res.Expenses)[0].ExtendProps.Amount)
	assert.Equal(s.T(), endOfMonthExpense.PaidAt.Format("2006-01-02"), (*res.Expenses)[1].Date.Format("2006-01-02"))
	assert.Equal(s.T(), "expense", (*res.Expenses)[0].ExtendProps.Type)
	assert.Equal(s.T(), endOfMonthExpense.Amount, (*res.Expenses)[1].ExtendProps.Amount)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_StatusUnauthorized() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	result := testutil.NewRequest().Get("/expenses").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestExpensesHandlerSuite) TestPostExpenses_Success_StatusOk() {
	user, cookieString := s.signIn()

	reqBody := api.StoreExpenseInput{
		PaidAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		Category: 1,
		Description: "description",
	}

	result := testutil.NewRequest().Post("/expenses").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.PostExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), reqBody.PaidAt.Format("2006-01-02"), res.Expense.PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), reqBody.Amount, res.Expense.Amount)
	assert.Equal(s.T(), reqBody.Category, res.Expense.Category)
	assert.Equal(s.T(), reqBody.Description, res.Expense.Description)

	// NOTE: バリデーションエラーがないことを確認
	assert.Equal(s.T(), api.StoreExpenseValidationError{}, res.Errors)

	var exists bool
	DBCon.Model(&models.Expense{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.True(s.T(), exists)
}

func (s *TestExpensesHandlerSuite) TestPostExpenses_ValidationErrorRequiredFields_StatusOk() {
	user, cookieString := s.signIn()

	reqBody := api.StoreExpenseInput{}

	result := testutil.NewRequest().Post("/expenses").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.PostExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	amountErrorMessages := []string{"金額は必須入力です。"}
	categoryErrorMessages := []string{"カテゴリは必須入力です。"}
	descriptionErrorMessages := []string{"適用は必須入力です。"}
	assert.ElementsMatch(s.T(), amountErrorMessages, *res.Errors.Amount)
	assert.ElementsMatch(s.T(), categoryErrorMessages, *res.Errors.Category)
	assert.ElementsMatch(s.T(), descriptionErrorMessages, *res.Errors.Description)

	var exists bool
	DBCon.Model(&models.Expense{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func (s *TestExpensesHandlerSuite) TestPostExpenses_ValidationErrorRequiredFields_StatusUnauthorized() {}

func TestExpensesHandler(t *testing.T) {
	// テストスイートを実施
	suite.Run(t, new(TestExpensesHandlerSuite))
}
