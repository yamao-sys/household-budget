package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/internal/validators"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type ExpenseService interface {
	FetchLists(userID int, beginningOfMonth *string) []models.Expense
	Create(userID int, requestParams *api.PostExpensesJSONRequestBody) (models.Expense, error)
	MappingValidationErrorStruct(err error) api.StoreExpenseValidationError
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

func (es *expenseService) Create(userID int, requestParams *api.PostExpensesJSONRequestBody) (models.Expense, error) {
	validationErr := validators.ValidateExpense(requestParams)
	if validationErr != nil {
		return models.Expense{}, validationErr
	}

	var expense models.Expense
	expense.UserID = userID
	expense.PaidAt = requestParams.PaidAt.Time
	expense.Amount = requestParams.Amount
	expense.Category = requestParams.Category
	expense.Description = requestParams.Description

	es.db.Create(&expense)

	return expense, nil
}

func (es *expenseService) MappingValidationErrorStruct(err error) api.StoreExpenseValidationError {
	var validationError api.StoreExpenseValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "paidAt":
				validationError.PaidAt = &messages
			case "amount":
				validationError.Amount = &messages
			case "category":
				validationError.Category = &messages
			case "description":
				validationError.Description = &messages
			}
		}
	}
	return validationError
}
