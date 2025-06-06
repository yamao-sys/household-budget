package services

import (
	api "apps/apis"
	"apps/internal/models"
	"apps/internal/validators"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type IncomeService interface {
	FetchLists(userID int, fromDate string, toDate string) []models.Income
	FetchTotalAmount(userID int, fromDate string, toDate string) []IncomeTotalAmount
	FetchClientTotalAmount(userID int, fromDate string, toDate string) []ClientTotalAmount
	Create(userID int, requestParams *api.PostIncomesJSONRequestBody) (models.Income, error)
	MappingValidationErrorStruct(err error) api.StoreIncomeValidationError
}

type incomeService struct {
	db *gorm.DB
}

type IncomeTotalAmount struct {
    ReceivedAt      time.Time
    TotalAmount int
}

type ClientTotalAmount struct {
	ClientName string
	TotalAmount int
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

func (is *incomeService) FetchTotalAmount(userID int, fromDate string, toDate string) []IncomeTotalAmount {
	var totalAmounts []IncomeTotalAmount

	is.db.Model(&models.Income{}).
		  Select("received_at, SUM(amount) AS total_amount").
		  Group("received_at").
		  Where("user_id = ? AND received_at BETWEEN ? AND ?", userID, fromDate, toDate).
		  Scan(&totalAmounts)
	return totalAmounts
}

func (is *incomeService) FetchClientTotalAmount(userID int, fromDate string, toDate string) []ClientTotalAmount {
	var clientTotalAmounts []ClientTotalAmount
	
	is.db.Model(&models.Income{}).
		  Select("client_name, SUM(amount) AS total_amount").
		  Group("client_name").
		  Where("user_id = ? AND received_at BETWEEN ? AND ?", userID, fromDate, toDate).
		  Scan(&clientTotalAmounts)
	return clientTotalAmounts
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
