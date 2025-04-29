package main

import (
	"apps/database"
	"apps/internal/models"
	"apps/test/factories"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	dbCon := database.Init()
	// NOTE: DBを閉じる
	defer func(cause error) {
		fmt.Println(cause)
		sqlDB, err := dbCon.DB()
		if err != nil {
			panic(err)
		}
		if cause = sqlDB.Close(); cause != nil {
			panic(cause)
		}
	}(nil)

	// NOTE: ログイン用ユーザの追加
	user := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_sign_in_1@example.com"}).(*models.User)
	dbCon.Create(&user)
	user2 := factories.UserFactory.MustCreateWithOption(map[string]interface{}{"Email": "test_sign_in_2@example.com"}).(*models.User)
	dbCon.Create(&user2)
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}
