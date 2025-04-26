package services

import (
	"apps/internal/models"
	"apps/test/factories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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

func TestExpenseService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestExpenseServiceSuite))
}
