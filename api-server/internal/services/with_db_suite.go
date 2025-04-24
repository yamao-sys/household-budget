package services

import (
	"apps/database"
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WithDBSuite struct {
	suite.Suite
}

var DBCon *gorm.DB
var ctx context.Context

// func (s *WithDBSuite) SetupSuite()                           {} // テストスイート実施前の処理
// func (s *WithDBSuite) TearDownSuite()                        {} // テストスイート終了後の処理
// func (s *WithDBSuite) SetupTest()                            {} // テストケース実施前の処理
// func (s *WithDBSuite) TearDownTest()                         {} // テストケース終了後の処理
// func (s *WithDBSuite) BeforeTest(suiteName, testName string) {} // テストケース実施前の処理
// func (s *WithDBSuite) AfterTest(suiteName, testName string)  {} // テストケース終了後の処理

func init() {
	txdb.Register("txdb-service", "mysql", database.GetDsn())
	ctx = context.Background()
}

func (s *WithDBSuite) SetDBCon() {
	db, _ := sql.Open("txdb-service", "connect")
	DBCon, _ = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
}

func (s *WithDBSuite) CloseDB() {
	db, _ := DBCon.DB()
	db.Close()
}
