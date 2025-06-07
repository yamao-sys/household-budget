package services

import (
	api "apps/apis"
	"apps/internal/models"
	"apps/test/factories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TestIncomeServiceSuite struct {
	WithDBSuite
}

var testIncomeService IncomeService

func (s *TestIncomeServiceSuite) SetupTest() {
	s.SetDBCon()

	testIncomeService = NewIncomeService(DBCon)
}

func (s *TestIncomeServiceSuite) TearDownTest() {
	s.CloseDB()
}

func (s *TestIncomeServiceSuite) TestIncomeFetchLists_WithFromDateAndToDate_Same() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	inRangeReceivedAtIncome1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local), "Amount": 10000}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome1)
	inRangeReceivedAtIncome2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local), "Amount": 10001}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome2)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherInRangeReceivedAtIncome)

	fromDate := "2025-04-01"
	toDate := "2025-04-01"
	fetchedIncomes := testIncomeService.FetchLists(user.ID,fromDate, toDate)

	// NOTE: 指定した日付の指定したユーザのIDに紐づく収入情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedIncomes))
	assert.Equal(s.T(), inRangeReceivedAtIncome1.ID, fetchedIncomes[0].ID)
	assert.Equal(s.T(), inRangeReceivedAtIncome2.ID, fetchedIncomes[1].ID)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchLists_WithFromDateAndToDate_Different() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 30, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	minInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minInRangeReceivedAtIncome)
	inRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome)
	maxInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxInRangeReceivedAtIncome)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangeReceivedAtIncomeIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherInRangeReceivedAtIncomeIncome)

	fromDate := "2025-03-31"
	toDate := "2025-04-02"
	fetchedIncomes := testIncomeService.FetchLists(user.ID,fromDate, toDate)

	// NOTE: 指定した期間の指定したユーザのIDに紐づく収入情報のみ取得されること
	assert.Equal(s.T(), 3, len(fetchedIncomes))
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ID, fetchedIncomes[0].ID)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ID, fetchedIncomes[1].ID)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ID, fetchedIncomes[2].ID)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchLists_WithFromDateAndWithoutToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	minInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minInRangeReceivedAtIncome)
	inRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherInRangeReceivedAtIncome)

	fromDate := "2025-04-01"
	fetchedIncomes := testIncomeService.FetchLists(user.ID, fromDate, "")

	// NOTE: 指定した期間の指定したユーザのIDに紐づく収入情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedIncomes))
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ID, fetchedIncomes[0].ID)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ID, fetchedIncomes[1].ID)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchLists_WithoutFromDateAndWithToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	inRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome)
	maxInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxInRangeReceivedAtIncome)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherInRangeReceivedAtIncome)

	toDate := "2025-04-01"
	fetchedIncomes := testIncomeService.FetchLists(user.ID, "", toDate)

	// NOTE: 指定した期間の指定したユーザのIDに紐づく収入情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedIncomes))
	assert.Equal(s.T(), inRangeReceivedAtIncome.ID, fetchedIncomes[0].ID)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ID, fetchedIncomes[1].ID)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchLists_WithoutFromDateAndToDate() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)
	Income1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&Income1)
	Income2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&Income2)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Income)
	DBCon.Create(&otherIncome)

	fetchedIncomes := testIncomeService.FetchLists(user.ID, "", "")

	// NOTE: 期間に関係なく指定したユーザのIDに紐づく収入情報のみ取得されること
	assert.Equal(s.T(), 2, len(fetchedIncomes))
	assert.Equal(s.T(), Income1.ID, fetchedIncomes[0].ID)
	assert.Equal(s.T(), Income2.ID, fetchedIncomes[1].ID)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchTotalAmount() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	inRangeReceivedAtIncomeFrom1_1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom1_1)
	inRangeReceivedAtIncomeFrom1_2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom1_2)
	inRangeReceivedAtIncomeFrom2_1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom2_1)
	inRangeReceivedAtIncomeFrom2_2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom2_2)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherBeginningOfMonthIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherBeginningOfMonthIncome)

	fromDate := "2025-04-01"
	toDate := "2025-04-02"
	totalAmounts := testIncomeService.FetchTotalAmount(user.ID, fromDate, toDate)

	// NOTE: 指定した期間で指定したユーザのIDに紐づく日ごとの支出合計のみ取得されること
	assert.Equal(s.T(), 2, len(totalAmounts))
	assert.Equal(s.T(), time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02"), totalAmounts[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncomeFrom1_1.Amount+inRangeReceivedAtIncomeFrom1_2.Amount, totalAmounts[0].TotalAmount)
	assert.Equal(s.T(), time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local).Format("2006-01-02"), totalAmounts[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncomeFrom2_1.Amount+inRangeReceivedAtIncomeFrom2_2.Amount, totalAmounts[1].TotalAmount)
}

func (s *TestIncomeServiceSuite) TestIncomeFetchClientTotalAmount() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ClientName": "テスト株式会社1", "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	inRangeReceivedAtIncomeFrom1_1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ClientName": "テスト株式会社1", "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom1_1)
	inRangeReceivedAtIncomeFrom1_2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ClientName": "テスト株式会社2", "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom1_2)
	inRangeReceivedAtIncomeFrom2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ClientName": "テスト株式会社1", "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncomeFrom2)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ClientName": "テスト株式会社2", "ReceivedAt": time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherBeginningOfMonthIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ClientName": "テスト株式会社1", "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherBeginningOfMonthIncome)

	fromDate := "2025-04-01"
	toDate := "2025-04-02"
	totalAmounts := testIncomeService.FetchClientTotalAmount(user.ID, fromDate, toDate)

	// NOTE: 指定した期間で指定したユーザのIDに紐づく顧客ごとの収入合計のみ取得されること
	assert.Equal(s.T(), 2, len(totalAmounts))
	assert.Equal(s.T(), "テスト株式会社1", totalAmounts[0].ClientName)
	assert.Equal(s.T(), inRangeReceivedAtIncomeFrom1_1.Amount+inRangeReceivedAtIncomeFrom2.Amount, totalAmounts[0].TotalAmount)
	assert.Equal(s.T(), "テスト株式会社2", totalAmounts[1].ClientName)
	assert.Equal(s.T(), inRangeReceivedAtIncomeFrom1_2.Amount, totalAmounts[1].TotalAmount)
}

func (s *TestIncomeServiceSuite) TestIncomeCreate_Success() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostIncomesJSONRequestBody{
		ReceivedAt: openapi_types.Date{Time: time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)},
		Amount: 10000,
		ClientName: "test client",
	}
	
	createdIncome, validationErr := testIncomeService.Create(user.ID, &requestParams)

	// NOTE: Incomeが作成されていることを確認
	assert.Equal(s.T(), user.ID, createdIncome.UserID)
	assert.Equal(s.T(), requestParams.ReceivedAt.Format("2006-01-02"), createdIncome.ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), requestParams.Amount, createdIncome.Amount)
	assert.Equal(s.T(), requestParams.ClientName, createdIncome.ClientName)

	// NOTE: バリデーションエラーがないことの確認
	assert.Nil(s.T(), validationErr)
	mappedValidationErr := testIncomeService.MappingValidationErrorStruct(validationErr)
	assert.Equal(s.T(), api.StoreIncomeValidationError{}, mappedValidationErr)

	// NOTE: DBに保存されていることを確認
	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ? AND received_at = ? AND amount = ? AND client_name = ?", user.ID, requestParams.ReceivedAt.Time, requestParams.Amount, requestParams.ClientName).Find(&exists)
	assert.True(s.T(), exists)
}

func (s *TestIncomeServiceSuite) TestIncomeCreate_ValidationErrorRequiredFields() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	requestParams := api.PostIncomesJSONRequestBody{}
	
	_, validationErr := testIncomeService.Create(user.ID, &requestParams)
	mappedValidationErr := testIncomeService.MappingValidationErrorStruct(validationErr)

	amountErrorMessages := []string{"金額は必須入力です。"}
	clientNameErrorMessages := []string{"顧客名は必須入力です。"}
	assert.ElementsMatch(s.T(), amountErrorMessages, *mappedValidationErr.Amount)
	assert.ElementsMatch(s.T(), clientNameErrorMessages, *mappedValidationErr.ClientName)

	// NOTE: DBに保存されていないことを確認
	var exists bool
	DBCon.Model(&models.Income{}).Select("count(*) > 0").Where("user_id = ?", user.ID).Find(&exists)
	assert.False(s.T(), exists)
}

func TestIncomeService(t *testing.T) {
	// テストスイートを実行
	suite.Run(t, new(TestIncomeServiceSuite))
}
