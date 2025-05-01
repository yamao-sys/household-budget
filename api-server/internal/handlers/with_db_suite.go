package handlers

import (
	api "apps/api"
	"apps/database"
	"apps/internal/middlewares"
	"apps/internal/models"
	"apps/internal/services"
	"apps/test/factories"
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/testutil"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WithDBSuite struct {
	suite.Suite
}

var (
	DBCon *gorm.DB
	// token string
	e *echo.Echo
	csrfToken string
	csrfTokenCookie string
)

// func (s *WithDBSuite) SetupSuite()                           {} // テストスイート実施前の処理
// func (s *WithDBSuite) TearDownSuite()                        {} // テストスイート終了後の処理
// func (s *WithDBSuite) SetupTest()                            {} // テストケース実施前の処理
// func (s *WithDBSuite) TearDownTest()                         {} // テストケース終了後の処理
// func (s *WithDBSuite) BeforeTest(suiteName, testName string) {} // テストケース実施前の処理
// func (s *WithDBSuite) AfterTest(suiteName, testName string)  {} // テストケース終了後の処理

func init() {
	txdb.Register("txdb-handler", "mysql", database.GetDsn())
	e = middlewares.ApplyMiddlewares(echo.New())
}

func (s *WithDBSuite) SetDBCon() {
	db, _ := sql.Open("txdb-handler", "connect")
	DBCon, _ = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
}

func (s *WithDBSuite) CloseDB() {
	db, _ := DBCon.DB()
	db.Close()
}

func (s *WithDBSuite) SetCsrfHeaderValues() {
	result := testutil.NewRequest().Get("/csrf").GoWithHTTPHandler(s.T(), e)

	var res api.GetCsrf200JSONResponse
	err := result.UnmarshalJsonToObject(&res)
	if err != nil {
		s.T().Error(err.Error())
	}

	csrfToken = res.CsrfToken
	csrfTokenCookie = result.Recorder.Result().Header.Values("Set-Cookie")[0]
}

func (s *WithDBSuite) initializeHandlers() {
	csrfServer := NewCsrfHandler()

	userService := services.NewUserService(DBCon)
	testUsersHandler := NewUsersHandler(userService)

	expenseService := services.NewExpenseService(DBCon)
	testExpensesHandler := NewExpensesHandler(expenseService)

	incomeService := services.NewIncomeService(DBCon)
	testIncomesHandler := NewIncomesHandler(incomeService)

	mainHandler := NewMainHandler(csrfServer, testUsersHandler, testExpensesHandler, testIncomesHandler)

	strictHandler := api.NewStrictHandler(mainHandler, []api.StrictMiddlewareFunc{middlewares.AuthMiddleware})
	api.RegisterHandlers(e, strictHandler)
}

func (s *WithDBSuite) signIn() (user *models.User, cookieString string) {
	// NOTE: テスト用ユーザの作成
	user = factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test@example.com"}).(*models.User)
	DBCon.Create(&user)

	reqBody := api.UserSignInInput{
		Email: "test@example.com",
		Password: "password",
	}
	result := testutil.NewRequest().Post("/users/signIn").WithHeader("Cookie", csrfTokenCookie).WithHeader(echo.HeaderXCSRFToken, csrfToken).WithJsonBody(reqBody).GoWithHTTPHandler(s.T(), e)
	cookieString = result.Recorder.Result().Header.Values("Set-Cookie")[0]

	return user, cookieString
}
