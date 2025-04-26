package services

import (
	"apps/internal/models"
	"time"

	"gorm.io/gorm"
)

type ExpenseService interface {
	FetchLists(userID int, beginningOfMonth *string) []models.Expense
}

type expenseService struct {
	db *gorm.DB
}

func NewExpenseService(db *gorm.DB) ExpenseService {
	return &expenseService{db}
}

func (es *expenseService) FetchLists(userID int, beginningOfMonth *string) []models.Expense {
	var expenses []models.Expense

	if beginningOfMonth == nil {
		es.db.Where("user_id = ?", userID).Find(&expenses)
		return expenses
	}
	// NOTE: 引数のmonthがある場合は月初と月末のBETWEENで検索する
	start, _ := time.Parse("2006-01-02", *beginningOfMonth)
	// NOTE: 月末を求める：翌月1日の前日
	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
	// NOTE: フォーマット（DBのDATE型に合わせて）
	startStr := start.Format("2006-01-02")
	endStr := end.Format("2006-01-02")

	es.db.Where("user_id = ? AND paid_at BETWEEN ? AND ?", userID, startStr, endStr).Find(&expenses)
	return expenses
}
