package handlers

import (
	api "apps/api"
	"apps/database"
	"apps/internal/middlewares"
	"apps/internal/services"
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

	mainHandler := NewMainHandler(csrfServer, testUsersHandler)

	strictHandler := api.NewStrictHandler(mainHandler, nil)
	api.RegisterHandlers(e, strictHandler)
}
