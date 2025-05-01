package handlers

import (
	"apps/api"
	"apps/internal/helpers"
	"apps/internal/services"
	"context"
	"strconv"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type IncomesHandler interface {
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

func (eh *incomesHandler) PostIncomes(ctx context.Context, request api.PostIncomesRequestObject) (api.PostIncomesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	createdIncome, validationErr := eh.incomeService.Create(userID, request.Body)

	resValidationError := eh.incomeService.MappingValidationErrorStruct(validationErr)

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
