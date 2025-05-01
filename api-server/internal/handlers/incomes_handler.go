package handlers

import (
	"apps/api"
	"apps/internal/helpers"
	"apps/internal/models"
	"apps/internal/services"
	"context"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type IncomesHandler interface {
	// Get Incomes
	// (GET /incomes)
	GetIncomes(ctx context.Context, request api.GetIncomesRequestObject) (api.GetIncomesResponseObject, error)
	// Post Income
	// (POST /incomes)
	PostIncomes(ctx context.Context, request api.PostIncomesRequestObject) (api.PostIncomesResponseObject, error)
}

type incomesHandler struct {
	incomeService services.IncomeService
}

func NewIncomesHandler(incomeService services.IncomeService) IncomesHandler {
	return &incomesHandler{incomeService}
}

func (ih *incomesHandler) GetIncomes(ctx context.Context, request api.GetIncomesRequestObject) (api.GetIncomesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	fromDate := request.Params.FromDate
	toDate := request.Params.ToDate

	var incomes []api.Income

	var fetchedIncomes []models.Income
	if fromDate != nil && toDate != nil {
		fetchedIncomes = ih.incomeService.FetchLists(userID, *fromDate, *toDate)
	} else if fromDate != nil && toDate == nil {
		fetchedIncomes = ih.incomeService.FetchLists(userID, *fromDate, "")
	} else if fromDate == nil && toDate != nil {
		fetchedIncomes = ih.incomeService.FetchLists(userID, "", *toDate)
	} else {
		fetchedIncomes = ih.incomeService.FetchLists(userID, "", "")
	}

	for _, income := range fetchedIncomes {
		incomes = append(incomes, api.Income{
			Amount: income.Amount,
			ClientName: income.ClientName,
			Id: strconv.Itoa(int(income.ID)),
			ReceivedAt: openapi_types.Date{Time: income.ReceivedAt},
		})
	}

	return api.GetIncomes200JSONResponse{FetchIncomeListsResponseJSONResponse: api.FetchIncomeListsResponseJSONResponse{Incomes: incomes}}, nil
}

func (ih *incomesHandler) PostIncomes(ctx context.Context, request api.PostIncomesRequestObject) (api.PostIncomesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	createdIncome, validationErr := ih.incomeService.Create(userID, request.Body)

	resValidationError := ih.incomeService.MappingValidationErrorStruct(validationErr)

	return api.PostIncomes200JSONResponse{StoreIncomeResponseJSONResponse: api.StoreIncomeResponseJSONResponse{
		Errors: resValidationError,
		Income: api.Income{
			Id: strconv.Itoa(createdIncome.ID),
			ReceivedAt: openapi_types.Date{Time: createdIncome.ReceivedAt},
			Amount: createdIncome.Amount,
			ClientName: createdIncome.ClientName,
		},
	}}, nil
}
