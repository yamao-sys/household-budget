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

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithFromDateAndToDate_Same() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&minOutOfRangePaidAtExpense)
	inRangePaidAtExpense1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense1)
	inRangePaidAtExpense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&inRangePaidAtExpense2)
	maxOutOfRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&maxOutOfRangePaidAtExpense)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangePaidAtExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "PaidAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Expense)
	DBCon.Create(&otherInRangePaidAtExpense)

	fromDate := "2025-04-01"
	toDate := "2025-04-01"
	fetchedExpenses := testExpenseService.FetchLists(user.ID,fromDate, toDate)

	// NOTE: 指定した日付の指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedExpenses))
	assert.Equal(s.T(), inRangePaidAtExpense1.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), inRangePaidAtExpense2.ID, fetchedExpenses[1].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithFromDateAndToDate_Different() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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

	fromDate := "2025-03-31"
	toDate := "2025-04-02"
	fetchedExpenses := testExpenseService.FetchLists(user.ID,fromDate, toDate)

	// NOTE: 指定した期間の指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 3, len(fetchedExpenses))
	assert.Equal(s.T(), minInRangePaidAtExpense.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), inRangePaidAtExpense.ID, fetchedExpenses[1].ID)
	assert.Equal(s.T(), maxInRangePaidAtExpense.ID, fetchedExpenses[2].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithFromDateAndWithoutToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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

	fromDate := "2025-04-01"
	fetchedExpenses := testExpenseService.FetchLists(user.ID, fromDate, "")

	// NOTE: 指定した期間の指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedExpenses))
	assert.Equal(s.T(), minInRangePaidAtExpense.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), inRangePaidAtExpense.ID, fetchedExpenses[1].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithoutFromDateAndWithToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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

	toDate := "2025-04-01"
	fetchedExpenses := testExpenseService.FetchLists(user.ID, "", toDate)

	// NOTE: 指定した期間の指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedExpenses))
	assert.Equal(s.T(), inRangePaidAtExpense.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), maxInRangePaidAtExpense.ID, fetchedExpenses[1].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchLists_WithoutFromDateAndToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)
	expense1 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Expense)
	DBCon.Create(&expense1)
	expense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Expense)
	DBCon.Create(&expense2)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Expense)
	DBCon.Create(&otherExpense)

	fetchedExpenses := testExpenseService.FetchLists(user.ID, "", "")

	// NOTE: 期間に関係なく指定したユーザのIDに紐づく支出情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedExpenses))
	assert.Equal(s.T(), expense1.ID, fetchedExpenses[0].ID)
	assert.Equal(s.T(), expense2.ID, fetchedExpenses[1].ID)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchTotalAmount() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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

	fromDate := "2025-04-01"
	toDate := "2025-04-02"
	totalAmounts := testExpenseService.FetchTotalAmount(user.ID, fromDate, toDate)

	// NOTE: 指定した期間で指定したユーザのIDに紐づく日ごとの支出合計のみ取得されること
	assert.Equal(s.T(), 2, len(totalAmounts))
	assert.Equal(s.T(), time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02"), totalAmounts[0].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_1.Amount+inRangePaidAtExpenseFrom1_2.Amount, totalAmounts[0].TotalAmount)
	assert.Equal(s.T(), time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local).Format("2006-01-02"), totalAmounts[1].PaidAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangePaidAtExpenseFrom2_1.Amount+inRangePaidAtExpenseFrom2_2.Amount, totalAmounts[1].TotalAmount)
}

func (s *TestExpenseServiceSuite) TestExpenseFetchCategoryTotalAmount() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

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

	fromDate := "2025-04-01"
	toDate := "2025-04-02"
	totalAmounts := testExpenseService.FetchCategoryTotalAmount(user.ID, fromDate, toDate)

	// NOTE: 指定した期間で指定したユーザのIDに紐づくカテゴリごとの支出合計のみ取得されること
	assert.Equal(s.T(), 2, len(totalAmounts))
	assert.Equal(s.T(), models.CategoryFood, totalAmounts[0].Category)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_1.Amount+inRangePaidAtExpenseFrom2.Amount, totalAmounts[0].TotalAmount)
	assert.Equal(s.T(), models.CategoryDailyGoods, totalAmounts[1].Category)
	assert.Equal(s.T(), inRangePaidAtExpenseFrom1_2.Amount, totalAmounts[1].TotalAmount)
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
	assert.Equal(s.T(), models.Category(requestParams.Category), createdExpense.Category)
	assert.Equal(s.T(), requestParams.Description, createdExpense.Description)

	// NOTE: バリデーションエラーがないことの確認
	assert.Nil(s.T(), validationErr)
	mappedValidationErr := testExpenseService.MappingValidationErrorStruct(validationErr)
	assert.Equal(s.T(), api.StoreExpenseValidationError{}, mappedValidationErr)

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
