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
