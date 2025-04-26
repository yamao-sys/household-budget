package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/test/factories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TestExpenseServiceSuite struct {
	WithDBSuite
}

var testExpenseService ExpenseService

func (s *TestExpenseServiceSuite) SetupTest() {
	s.SetDBCon()

	testExpenseService = NewExpenseService(DBCon)
}

func (s *TestExpenseServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithoutBeginngOfMonth() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)
	expense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Expense)
	DBCon.Create(&expense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Expense)
	DBCon.Create(&otherExpense)

	fetchedExpenses := testExpenseService.FetchLists(user.ID, nil)

	// NOTE: 指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 1, len(fetchedExpenses))
	assert.Equal(s.T(), expense.ID, fetchedExpenses[0].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithBeginngOfMonth() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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
	fetchedExpenses := testExpenseService.FetchLists(user.ID, &beginningOfMonth)

	// NOTE: 指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedExpenses))
	assert.Equal(s.T(), beginningOfMonthExpense.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), endOfMonthExpense.ID, fetchedExpenses[1].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseCreate_Success() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostExpensesJSONRequestBody{
		PaidAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		Category: 1,
		Description: "description",
	}
	
	createdExpense, validationErr := testExpenseService.Create(user.ID, &requestParams)

	// NOTE: Expenseが作成されていることを確認
	assert.Equal(s.T(), user.ID, createdExpense.UserID)
	assert.Equal(s.T(), requestParams.PaidAt.Format("2006-01-02"), createdExpense.PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), requestParams.Amount, createdExpense.Amount)
	assert.Equal(s.T(), requestParams.Category, createdExpense.Category)
	assert.Equal(s.T(), requestParams.Description, createdExpense.Description)

	// NOTE: バリデーションエラーがないことの確認
	assert.Nil(s.T(), validationErr)

	// NOTE: DBに保存されていることを確認
	var exists bool
	DBCon.Model(&models.Expense{}).Select("count(*) > 0").Where("user_id = ? AND paid_at = ? AND amount = ? AND category = ? AND description = ?", user.ID, requestParams.PaidAt.Time, requestParams.Amount, requestParams.Category, requestParams.Description).Find(&exists)
	assert.True(s.T(), exists)
}

func (s *TestExpenseServiceSuite) TestExpenseCreate_ValidationErrorRequiredFields() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostExpensesJSONRequestBody{}
	
	_, validationErr := testExpenseService.Create(user.ID, &requestParams)
	mappedValidationErr := testExpenseService.MappingValidationErrorStruct(validationErr)

	amountErrorMessages := []string{"金額は必須入力です。"}
	categoryErrorMessages := []string{"カテゴリは必須入力です。"}
	descriptionErrorMessages := []string{"適用は必須入力です。"}
	assert.ElementsMatch(s.T(), amountErrorMessages, *mappedValidationErr.Amount)
	assert.ElementsMatch(s.T(), categoryErrorMessages, *mappedValidationErr.Category)
	assert.ElementsMatch(s.T(), descriptionErrorMessages, *mappedValidationErr.Description)

	// NOTE: DBに保存されていないことを確認
	var exists bool
	DBCon.Model(&models.Expense{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func TestExpenseService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestExpenseServiceSuite))
}
