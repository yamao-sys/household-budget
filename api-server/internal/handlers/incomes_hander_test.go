package handlers

import (
	api "apps/api"
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

func (s *TestIncomesHandlerSuite) TestGetIncomes_WithFromDateAndToDate_Same_StatusOk() {
	user, cookieString := s.signIn()

	minOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 3, 31, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&minOutOfRangeReceivedAtIncome)
	inRangeReceivedAtIncome1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome1)
	inRangeReceivedAtIncome2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&inRangeReceivedAtIncome2)
	maxOutOfRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&maxOutOfRangeReceivedAtIncome)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherInRangeReceivedAtIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser, "ReceivedAt": time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local)}).(*models.Income)
	DBCon.Create(&otherInRangeReceivedAtIncome)

	result := testutil.NewRequest().Get("/incomes?fromDate=2025-04-01&toDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Incomes))
	assert.Equal(s.T(), strconv.Itoa(inRangeReceivedAtIncome1.ID), res.Incomes[0].Id)
	assert.Equal(s.T(), inRangeReceivedAtIncome1.ReceivedAt.Format("2006-01-02"), res.Incomes[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncome1.Amount, res.Incomes[0].Amount)
	assert.Equal(s.T(), inRangeReceivedAtIncome1.ClientName, res.Incomes[0].ClientName)

	assert.Equal(s.T(), strconv.Itoa(inRangeReceivedAtIncome2.ID), res.Incomes[1].Id)
	assert.Equal(s.T(), inRangeReceivedAtIncome2.ReceivedAt.Format("2006-01-02"), res.Incomes[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncome2.Amount, res.Incomes[1].Amount)
	assert.Equal(s.T(), inRangeReceivedAtIncome2.ClientName, res.Incomes[1].ClientName)
}

func (s *TestIncomesHandlerSuite) TestGetIncomes_WithFromDateAndToDate_Different_StatusOk() {
	user, cookieString := s.signIn()

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

	result := testutil.NewRequest().Get("/incomes?fromDate=2025-03-31&toDate=2025-04-02").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 3, len(res.Incomes))
	assert.Equal(s.T(), strconv.Itoa(minInRangeReceivedAtIncome.ID), res.Incomes[0].Id)
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), minInRangeReceivedAtIncome.Amount, res.Incomes[0].Amount)
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ClientName, res.Incomes[0].ClientName)

	assert.Equal(s.T(), strconv.Itoa(inRangeReceivedAtIncome.ID), res.Incomes[1].Id)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncome.Amount, res.Incomes[1].Amount)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ClientName, res.Incomes[1].ClientName)

	assert.Equal(s.T(), strconv.Itoa(maxInRangeReceivedAtIncome.ID), res.Incomes[2].Id)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[2].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.Amount, res.Incomes[2].Amount)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ClientName, res.Incomes[2].ClientName)
}

func (s *TestIncomesHandlerSuite) TestGetIncomes_WithFromDateAndWithoutToDate_StatusOk() {
	user, cookieString := s.signIn()

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

	result := testutil.NewRequest().Get("/incomes?fromDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Incomes))
	assert.Equal(s.T(), strconv.Itoa(minInRangeReceivedAtIncome.ID), res.Incomes[0].Id)
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), minInRangeReceivedAtIncome.Amount, res.Incomes[0].Amount)
	assert.Equal(s.T(), minInRangeReceivedAtIncome.ClientName, res.Incomes[0].ClientName)

	assert.Equal(s.T(), strconv.Itoa(inRangeReceivedAtIncome.ID), res.Incomes[1].Id)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncome.Amount, res.Incomes[1].Amount)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ClientName, res.Incomes[1].ClientName)
}

func (s *TestIncomesHandlerSuite) TestGetIncomes_WithoutFromDateAndWithToDate_StatusOk() {
	user, cookieString := s.signIn()

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

	result := testutil.NewRequest().Get("/incomes?toDate=2025-04-01").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Incomes))
	assert.Equal(s.T(), strconv.Itoa(inRangeReceivedAtIncome.ID), res.Incomes[0].Id)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), inRangeReceivedAtIncome.Amount, res.Incomes[0].Amount)
	assert.Equal(s.T(), inRangeReceivedAtIncome.ClientName, res.Incomes[0].ClientName)

	assert.Equal(s.T(), strconv.Itoa(maxInRangeReceivedAtIncome.ID), res.Incomes[1].Id)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ReceivedAt.Format("2006-01-02"), res.Incomes[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.Amount, res.Incomes[1].Amount)
	assert.Equal(s.T(), maxInRangeReceivedAtIncome.ClientName, res.Incomes[1].ClientName)
}

func (s *TestIncomesHandlerSuite) TestGetIncomes_WithoutFromDateAndToDate_StatusOk() {
	user, cookieString := s.signIn()

	expense1 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Income)
	DBCon.Create(&expense1)
	expense2 := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user}).(*models.Income)
	DBCon.Create(&expense2)

	otherUser := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_other@example.com"}).(*models.User)
	DBCon.Create(&otherUser)
	otherIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *otherUser}).(*models.Income)
	DBCon.Create(&otherIncome)

	result := testutil.NewRequest().Get("/incomes").WithHeader("Cookie", csrfTokenCookie+"; "+cookieString).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusOK, result.Code())

	var res api.GetIncomes200JSONResponse
	err := result.UnmarshalBodyToObject(&res)
	assert.NoError(s.T(), err, "error unmarshaling response")

	assert.Equal(s.T(), 2, len(res.Incomes))
	assert.Equal(s.T(), strconv.Itoa(expense1.ID), res.Incomes[0].Id)
	assert.Equal(s.T(), expense1.ReceivedAt.Format("2006-01-02"), res.Incomes[0].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), expense1.Amount, res.Incomes[0].Amount)
	assert.Equal(s.T(), expense1.ClientName, res.Incomes[0].ClientName)

	assert.Equal(s.T(), strconv.Itoa(expense2.ID), res.Incomes[1].Id)
	assert.Equal(s.T(), expense2.ReceivedAt.Format("2006-01-02"), res.Incomes[1].ReceivedAt.Format("2006-01-02"))
	assert.Equal(s.T(), expense2.Amount, res.Incomes[1].Amount)
	assert.Equal(s.T(), expense2.ClientName, res.Incomes[1].ClientName)
}

func (s *TestIncomesHandlerSuite) TestGetIncomes_StatusUnauthorized() {
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	result := testutil.NewRequest().Get("/incomes").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).GoWithHTTPHandler(s.T(), e)
	assert.Equal(s.T(), http.StatusUnauthorized, result.Code())
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
