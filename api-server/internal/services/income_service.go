package services

import (
	"apps/api"
	"apps/internal/models"
	"apps/internal/validators"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type IncomeService interface {
	Create(userID int, requestParams *api.PostIncomesJSONRequestBody) (models.Income, error)
	MappingValidationErrorStruct(err error) api.StoreIncomeValidationError
}

type incomeService struct {
	db *gorm.DB
}

func NewIncomeService(db *gorm.DB) IncomeService {
	return &incomeService{db}
}

func (es *incomeService) Create(userID int, requestParams *api.PostIncomesJSONRequestBody) (models.Income, error) {
	validationErr := validators.ValidateIncome(requestParams)
	if validationErr != nil {
		return models.Income{}, validationErr
	}

	var income models.Income
	income.UserID = userID
	income.ReceivedAt = requestParams.ReceivedAt.Time
	income.Amount = requestParams.Amount
	income.ClientName = requestParams.ClientName

	es.db.Create(&income)

	return income, nil
}

func (es *incomeService) MappingValidationErrorStruct(err error) api.StoreIncomeValidationError {
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
