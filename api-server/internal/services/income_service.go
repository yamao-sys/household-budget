package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/internal/validators"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type IncomeService interface {
	FetchLists(userID int, fromDate string, toDate string) []models.Income
	Create(userID int, requestParams *api.PostIncomesJSONRequestBody) (models.Income, error)
	MappingValidationErrorStruct(err error) api.StoreIncomeValidationError
}

type incomeService struct {
	db *gorm.DB
}

func NewIncomeService(db *gorm.DB) IncomeService {
	return &incomeService{db}
}

func (is *incomeService) FetchLists(userID int, fromDate string, toDate string) []models.Income {
	var incomes []models.Income

	if fromDate != "" && toDate != "" {
		if fromDate == toDate {
			date := fromDate
			is.db.Where("user_id = ? AND received_at = ?", userID, date).Find(&incomes)
			return incomes
		}

		is.db.Where("user_id = ? AND received_at BETWEEN ? AND ?", userID, fromDate, toDate).Find(&incomes)
	} else if fromDate != "" && toDate == "" {
		is.db.Where("user_id = ? AND received_at >= ?", userID, fromDate).Find(&incomes)
	} else if fromDate == "" && toDate != "" {
		is.db.Where("user_id = ? AND received_at <= ?", userID, toDate).Find(&incomes)
	} else {
		is.db.Where("user_id = ?", userID).Find(&incomes)
	}

	return incomes
}

func (is *incomeService) Create(userID int, requestParams *api.PostIncomesJSONRequestBody) (models.Income, error) {
	validationErr := validators.ValidateIncome(requestParams)
	if validationErr != nil {
		return models.Income{}, validationErr
	}

	var income models.Income
	income.UserID = userID
	income.ReceivedAt = requestParams.ReceivedAt.Time
	income.Amount = requestParams.Amount
	income.ClientName = requestParams.ClientName

	is.db.Create(&income)

	return income, nil
}

func (is *incomeService) MappingValidationErrorStruct(err error) api.StoreIncomeValidationError {
	var validationError api.StoreIncomeValidationError
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
			case "clientName":
				validationError.ClientName = &messages
			}
		}
	}
	return validationError
}
