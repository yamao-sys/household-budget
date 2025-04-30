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
	FetchLists(userID int, fromDate string, toDate string) []models.Expense
	FetchTotalAmount(userID int, fromDate string, toDate string) []TotalAmount
	FetchCategoryTotalAmount(userID int, fromDate string, toDate string) []CategoryTotalAmount
	Create(userID int, requestParams *api.PostExpensesJSONRequestBody) (models.Expense, error)
	MappingValidationErrorStruct(err error) api.StoreExpenseValidationError
}

type expenseService struct {
	db *gorm.DB
}

type TotalAmount struct {
	PaidAt time.Time
	TotalAmount int
}

type CategoryTotalAmount struct {
	Category models.Category
	TotalAmount int
}

func NewExpenseService(db *gorm.DB) ExpenseService {
	return &expenseService{db}
}

func (es *expenseService) FetchLists(userID int, fromDate string, toDate string) []models.Expense {
	var expenses []models.Expense

	if fromDate != "" && toDate != "" {
		if fromDate == toDate {
			date := fromDate
			es.db.Where("user_id = ? AND paid_at = ?", userID, date).Find(&expenses)
			return expenses
		}

		es.db.Where("user_id = ? AND paid_at BETWEEN ? AND ?", userID, fromDate, toDate).Find(&expenses)
	} else if fromDate != "" && toDate == "" {
		es.db.Where("user_id = ? AND paid_at >= ?", userID, fromDate).Find(&expenses)
	} else if fromDate == "" && toDate != "" {
		es.db.Where("user_id = ? AND paid_at <= ?", userID, toDate).Find(&expenses)
	} else {
		es.db.Where("user_id = ?", userID).Find(&expenses)
	}

	return expenses
}

func (es *expenseService) FetchTotalAmount(userID int, fromDate string, toDate string) []TotalAmount {
	var totalAmounts []TotalAmount

	es.db.Model(&models.Expense{}).
		  Select("paid_at, SUM(amount) AS total_amount").
		  Group("paid_at").
		  Where("user_id = ? AND paid_at BETWEEN ? AND ?", userID, fromDate, toDate).
		  Scan(&totalAmounts)
	return totalAmounts
}

func (es *expenseService) FetchCategoryTotalAmount(userID int, fromDate string, toDate string) []CategoryTotalAmount {
	var categoryTotalAmounts []CategoryTotalAmount
	
	es.db.Model(&models.Expense{}).
		  Select("category, SUM(amount) AS total_amount").
		  Group("category").
		  Where("user_id = ? AND paid_at BETWEEN ? AND ?", userID, fromDate, toDate).
		  Scan(&categoryTotalAmounts)
	return categoryTotalAmounts
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
	expense.Category = models.Category(requestParams.Category)
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
