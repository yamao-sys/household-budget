package handlers

import (
	api "apps/apis"
	"apps/internal/models"
	"apps/test/factories"
	"net/http"
	"strconv"
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

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithFromDateAndToDate_Same_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	inRangePaidAtExpense1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local), "Amount": 10000}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense1)
	inRangePaidAtExpense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local), "Amount": 10001}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense2)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherInRangePaidAtExpense)

	result := testutil.NewRequest().Get("/expenses?fromDate=2025-04-01&toDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Expenses))
	assert.Equal(s.T(), strconv.Itoa(inRangePaidAtExpense1.ID), res.Expenses[0].Id)
	assert.Equal(s.T(), inRangePaidAtExpense1.PaidAt.Format("2006-01-02"), res.Expenses[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpense1.Amount, res.Expenses[0].Amount)
	assert.Equal(s.T(), int(inRangePaidAtExpense1.Category), res.Expenses[0].Category)
	assert.Equal(s.T(), inRangePaidAtExpense1.Description, res.Expenses[0].Description)

	assert.Equal(s.T(), strconv.Itoa(inRangePaidAtExpense2.ID), res.Expenses[1].Id)
	assert.Equal(s.T(), inRangePaidAtExpense2.PaidAt.Format("2006-01-02"), res.Expenses[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpense2.Amount, res.Expenses[1].Amount)
	assert.Equal(s.T(), int(inRangePaidAtExpense2.Category), res.Expenses[1].Category)
	assert.Equal(s.T(), inRangePaidAtExpense2.Description, res.Expenses[1].Description)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithFromDateAndToDate_Different_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 30, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	minInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minInRangePaidAtExpense)
	inRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense)
	maxInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxInRangePaidAtExpense)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangePaidAtExpenseExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherInRangePaidAtExpenseExpense)

	result := testutil.NewRequest().Get("/expenses?fromDate=2025-03-31&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 3, len(res.Expenses))
	assert.Equal(s.T(), strconv.Itoa(minInRangePaidAtExpense.ID), res.Expenses[0].Id)
	assert.Equal(s.T(), minInRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), minInRangePaidAtExpense.Amount, res.Expenses[0].Amount)
	assert.Equal(s.T(), int(minInRangePaidAtExpense.Category), res.Expenses[0].Category)
	assert.Equal(s.T(), minInRangePaidAtExpense.Description, res.Expenses[0].Description)

	assert.Equal(s.T(), strconv.Itoa(inRangePaidAtExpense.ID), res.Expenses[1].Id)
	assert.Equal(s.T(), inRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpense.Amount, res.Expenses[1].Amount)
	assert.Equal(s.T(), int(inRangePaidAtExpense.Category), res.Expenses[1].Category)
	assert.Equal(s.T(), inRangePaidAtExpense.Description, res.Expenses[1].Description)

	assert.Equal(s.T(), strconv.Itoa(maxInRangePaidAtExpense.ID), res.Expenses[2].Id)
	assert.Equal(s.T(), maxInRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[2].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), maxInRangePaidAtExpense.Amount, res.Expenses[2].Amount)
	assert.Equal(s.T(), int(maxInRangePaidAtExpense.Category), res.Expenses[2].Category)
	assert.Equal(s.T(), maxInRangePaidAtExpense.Description, res.Expenses[2].Description)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithFromDateAndWithoutToDate_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	minInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minInRangePaidAtExpense)
	inRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherInRangePaidAtExpense)

	result := testutil.NewRequest().Get("/expenses?fromDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Expenses))
	assert.Equal(s.T(), strconv.Itoa(minInRangePaidAtExpense.ID), res.Expenses[0].Id)
	assert.Equal(s.T(), minInRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), minInRangePaidAtExpense.Amount, res.Expenses[0].Amount)
	assert.Equal(s.T(), int(minInRangePaidAtExpense.Category), res.Expenses[0].Category)
	assert.Equal(s.T(), minInRangePaidAtExpense.Description, res.Expenses[0].Description)

	assert.Equal(s.T(), strconv.Itoa(inRangePaidAtExpense.ID), res.Expenses[1].Id)
	assert.Equal(s.T(), inRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpense.Amount, res.Expenses[1].Amount)
	assert.Equal(s.T(), int(inRangePaidAtExpense.Category), res.Expenses[1].Category)
	assert.Equal(s.T(), inRangePaidAtExpense.Description, res.Expenses[1].Description)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithoutFromDateAndWithToDate_StatusOk() {
	user, cookieString := s.signIn()

	inRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense)
	maxInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxInRangePaidAtExpense)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherInRangePaidAtExpense)

	result := testutil.NewRequest().Get("/expenses?toDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Expenses))
	assert.Equal(s.T(), strconv.Itoa(inRangePaidAtExpense.ID), res.Expenses[0].Id)
	assert.Equal(s.T(), inRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpense.Amount, res.Expenses[0].Amount)
	assert.Equal(s.T(), int(inRangePaidAtExpense.Category), res.Expenses[0].Category)
	assert.Equal(s.T(), inRangePaidAtExpense.Description, res.Expenses[0].Description)

	assert.Equal(s.T(), strconv.Itoa(maxInRangePaidAtExpense.ID), res.Expenses[1].Id)
	assert.Equal(s.T(), maxInRangePaidAtExpense.PaidAt.Format("2006-01-02"), res.Expenses[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), maxInRangePaidAtExpense.Amount, res.Expenses[1].Amount)
	assert.Equal(s.T(), int(maxInRangePaidAtExpense.Category), res.Expenses[1].Category)
	assert.Equal(s.T(), maxInRangePaidAtExpense.Description, res.Expenses[1].Description)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_WithoutFromDateAndToDate_StatusOk() {
	user, cookieString := s.signIn()

	expense1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&expense1)
	expense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&expense2)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Expense)
	DBCon.Create(&otherExpense)

	result := testutil.NewRequest().Get("/expenses").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpenses200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Expenses))
	assert.Equal(s.T(), strconv.Itoa(expense1.ID), res.Expenses[0].Id)
	assert.Equal(s.T(), expense1.PaidAt.Format("2006-01-02"), res.Expenses[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), expense1.Amount, res.Expenses[0].Amount)
	assert.Equal(s.T(), int(expense1.Category), res.Expenses[0].Category)
	assert.Equal(s.T(), expense1.Description, res.Expenses[0].Description)

	assert.Equal(s.T(), strconv.Itoa(expense2.ID), res.Expenses[1].Id)
	assert.Equal(s.T(), expense2.PaidAt.Format("2006-01-02"), res.Expenses[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), expense2.Amount, res.Expenses[1].Amount)
	assert.Equal(s.T(), int(expense2.Category), res.Expenses[1].Category)
	assert.Equal(s.T(), expense2.Description, res.Expenses[1].Description)
}

func (s *TestExpensesHandlerSuite) TestGetExpenses_StatusUnauthorized() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	result := testutil.NewRequest().Get("/expenses").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestExpensesHandlerSuite) TestGetExpensesTotalAmounts_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	inRangePaidAtExpenseFrom1_1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom1_1)
	inRangePaidAtExpenseFrom1_2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom1_2)
	inRangePaidAtExpenseFrom2_1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom2_1)
	inRangePaidAtExpenseFrom2_2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom2_2)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherBeginningOfMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherBeginningOfMonthExpense)

	result := testutil.NewRequest().Get("/expenses/totalAmounts?fromDate=2025-04-01&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpensesTotalAmounts200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.TotalAmounts))
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_1.PaidAt.Format("2006-01-02"), res.TotalAmounts[0].Date.Format("2006-01-02"))
	assert.Equal(s.T(), "expense", res.TotalAmounts[0].ExtendProps.Type)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_1.Amount+inRangePaidAtExpenseFrom1_2.Amount, res.TotalAmounts[0].ExtendProps.TotalAmount)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom2_1.PaidAt.Format("2006-01-02"), res.TotalAmounts[1].Date.Format("2006-01-02"))
	assert.Equal(s.T(), "expense", res.TotalAmounts[1].ExtendProps.Type)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom2_1.Amount+inRangePaidAtExpenseFrom2_2.Amount, res.TotalAmounts[1].ExtendProps.TotalAmount)
}

func (s *TestExpensesHandlerSuite) TestGetExpensesTotalAmounts_StatusUnauthorized() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	result := testutil.NewRequest().Get("/expenses/totalAmounts?fromDate=2025-04-01&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
}

func (s *TestExpensesHandlerSuite) TestGetExpensesCategoryTotalAmounts_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "Category": models.CategoryFood, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	inRangePaidAtExpenseFrom1_1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "Category": models.CategoryFood, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom1_1)
	inRangePaidAtExpenseFrom1_2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "Category": models.CategoryDailyGoods, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom1_2)
	inRangePaidAtExpenseFrom2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "Category": models.CategoryFood, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpenseFrom2)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "Category": models.CategoryDailyGoods, "PaidAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherBeginningOfMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "Category": models.CategoryFood, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherBeginningOfMonthExpense)

	result := testutil.NewRequest().Get("/expenses/categoryTotalAmounts?fromDate=2025-04-01&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetExpensesCategoryTotalAmounts200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.TotalAmounts))
	assert.Equal(s.T(), int(models.CategoryFood), res.TotalAmounts[0].Category)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_1.Amount+inRangePaidAtExpenseFrom2.Amount, res.TotalAmounts[0].TotalAmount)
	assert.Equal(s.T(), int(models.CategoryDailyGoods), res.TotalAmounts[1].Category)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_2.Amount, res.TotalAmounts[1].TotalAmount)
}

func (s *TestExpensesHandlerSuite) TestGetExpensesCategoryTotalAmounts_StatusUnauthorized() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	result := testutil.NewRequest().Get("/expenses/categoryTotalAmounts?fromDate=2025-04-01&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
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
