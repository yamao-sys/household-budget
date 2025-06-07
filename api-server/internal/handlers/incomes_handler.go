package handlers

import (
	api "apps/apis"
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
	// Get Incomes TotalAmounts
	// (GET /incomes/totalAmounts)
	GetIncomesTotalAmounts(ctx context.Context, request api.GetIncomesTotalAmountsRequestObject) (api.GetIncomesTotalAmountsResponseObject, error)
	// Get Incomes ClientTotalAmounts
	// (GET /incomes/clientTotalAmounts)
	GetIncomesClientTotalAmounts(ctx context.Context, request api.GetIncomesClientTotalAmountsRequestObject) (api.GetIncomesClientTotalAmountsResponseObject, error)
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
			Id: strconv.Itoa(income.ID),
			ReceivedAt: openapi_types.Date{Time: income.ReceivedAt},
		})
	}

	return api.GetIncomes200JSONResponse(api.IncomeLists{Incomes: incomes}), nil
}

func (ih *incomesHandler) GetIncomesTotalAmounts(ctx context.Context, request api.GetIncomesTotalAmountsRequestObject) (api.GetIncomesTotalAmountsResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	fromDate := request.Params.FromDate
	toDate := request.Params.ToDate

	totalAmounts := ih.incomeService.FetchTotalAmount(userID, fromDate, toDate)

	var resTotalAmounts []api.TotalAmountLists
	for _, totalAmount := range totalAmounts {
		resTotalAmounts = append(resTotalAmounts, api.TotalAmountLists{
			Date: openapi_types.Date{Time: totalAmount.ReceivedAt},
			ExtendProps: struct{TotalAmount int "json:\"totalAmount\""; Type string "json:\"type\""} {
				TotalAmount: totalAmount.TotalAmount,
				Type: "income",
			},
		})
	}
	return api.GetIncomesTotalAmounts200JSONResponse(api.TotalAmountListsResponse{TotalAmounts: resTotalAmounts}), nil
}

func (ih *incomesHandler) GetIncomesClientTotalAmounts(ctx context.Context, request api.GetIncomesClientTotalAmountsRequestObject) (api.GetIncomesClientTotalAmountsResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	fromDate := request.Params.FromDate
	toDate := request.Params.ToDate

	categoryTotalAmounts := ih.incomeService.FetchClientTotalAmount(userID, fromDate, toDate)

	var resClientTotalAmounts []api.ClientTotalAmountLists
	for _, categoryTotalAmount := range categoryTotalAmounts {
		resClientTotalAmounts = append(resClientTotalAmounts, api.ClientTotalAmountLists{
			ClientName: categoryTotalAmount.ClientName,
			TotalAmount: categoryTotalAmount.TotalAmount,
		})
	}
	return api.GetIncomesClientTotalAmounts200JSONResponse(api.ClientTotalAmountListsResponse{TotalAmounts: resClientTotalAmounts}), nil
}

func (ih *incomesHandler) PostIncomes(ctx context.Context, request api.PostIncomesRequestObject) (api.PostIncomesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	createdIncome, validationErr := ih.incomeService.Create(userID, request.Body)

	resValidationError := ih.incomeService.MappingValidationErrorStruct(validationErr)

	return api.PostIncomes200JSONResponse(api.StoreIncomeResponse{
		Errors: resValidationError,
		Income: api.Income{
			Id: strconv.Itoa(createdIncome.ID),
			ReceivedAt: openapi_types.Date{Time: createdIncome.ReceivedAt},
			Amount: createdIncome.Amount,
			ClientName: createdIncome.ClientName,
		},
	}), nil
}
