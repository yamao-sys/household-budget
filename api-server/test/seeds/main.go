package main

import (
	"apps/database"
	"apps/internal/models"
	"apps/test/factories"
	"fmt"
	"os"
	"time"

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

	// NOTE: それぞれのユーザに対して、10日と20日の支出レコードを作成
	tenth := time.Date(2025, 5, 10, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	twentieth := time.Date(2025, 5, 20, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))

	tenthOfLastMonth := time.Date(2025, 4, 10, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	twentiethOfLastMonth := time.Date(2025, 4, 20, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	
	tenthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": tenth, "Amount": 10000, "Category": models.CategoryFood}).(*models.Expense)
	dbCon.Create(&tenthExpense)
	tenthExpense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": tenth, "Amount": 5000, "Category": models.CategoryDailyGoods}).(*models.Expense)
	dbCon.Create(&tenthExpense2)
	twentiethExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "PaidAt": twentieth, "Amount": 20000, "Category": models.CategoryEntertainment}).(*models.Expense)
	dbCon.Create(&twentiethExpense)
	twentiethIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user, "ReceivedAt": twentieth, "Amount": 1100000, "ClientName": "テスト株式会社1"}).(*models.Income)
	dbCon.Create(&twentiethIncome)

	tenthOfLastMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user2, "PaidAt": tenthOfLastMonth, "Amount": 10000, "Category": models.CategoryFood}).(*models.Expense)
	dbCon.Create(&tenthOfLastMonthExpense)
	tenthOfLastMonthExpense2 := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user2, "PaidAt": tenthOfLastMonth, "Amount": 5000, "Category": models.CategoryDailyGoods}).(*models.Expense)
	dbCon.Create(&tenthOfLastMonthExpense2)
	twentiethOfLastMonthExpense := factories.ExpenseFactory.MustCreateWithOption(map[string]interface{}{"User": *user2, "PaidAt": twentiethOfLastMonth, "Amount": 20000, "Category": models.CategoryEntertainment}).(*models.Expense)
	dbCon.Create(&twentiethOfLastMonthExpense)
	twentiethOfLastMonthIncome := factories.IncomeFactory.MustCreateWithOption(map[string]interface{}{"User": *user2, "ReceivedAt": twentiethOfLastMonth, "Amount": 1100000, "ClientName": "テスト株式会社2"}).(*models.Income)
	dbCon.Create(&twentiethOfLastMonthIncome)
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}
