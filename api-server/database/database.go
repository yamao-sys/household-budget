package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() *gorm.DB {
	// DBインスタンス生成
	DB, err = gorm.Open(mysql.Open(GetDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func Close(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

func GetDsn() string {
	baseDsn := os.Getenv("MYSQL_USER") +
				":" + os.Getenv("MYSQL_PASS") +
				"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" +
				os.Getenv("MYSQL_DBNAME") +
				"?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true&loc=Local"
	
	if os.Getenv("APP_ENV") != "production" {
		return baseDsn
	}

	// NOTE: 本番環境ではTiDBに接続するためにTLSを有効にする
	return baseDsn+"&tls=true"
}
