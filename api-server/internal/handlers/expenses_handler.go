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

type ExpensesHandler interface {
	// Get Expenses
	// (GET /expenses)
	GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error)
	// Get Expenses TotalAmounts
	// (GET /expenses/totalAmounts)
	GetExpensesTotalAmounts(ctx context.Context, request api.GetExpensesTotalAmountsRequestObject) (api.GetExpensesTotalAmountsResponseObject, error)
	// Post Expense
	// (POST /expenses)
	PostExpenses(ctx context.Context, request api.PostExpensesRequestObject) (api.PostExpensesResponseObject, error)
}

type expensesHandler struct {
	expenseService services.ExpenseService
}

func NewExpensesHandler(expenseService services.ExpenseService) ExpensesHandler {
	return &expensesHandler{expenseService}
}

func (eh *expensesHandler) GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	fromDate := request.Params.FromDate
	toDate := request.Params.ToDate

	var expenses []api.Expense

	var fetchedExpenses []models.Expense
	if fromDate != nil && toDate != nil {
		fetchedExpenses = eh.expenseService.FetchLists(userID, *fromDate, *toDate)
	} else if fromDate != nil && toDate == nil {
		fetchedExpenses = eh.expenseService.FetchLists(userID, *fromDate, "")
	} else if fromDate == nil && toDate != nil {
		fetchedExpenses = eh.expenseService.FetchLists(userID, "", *toDate)
	} else {
		fetchedExpenses = eh.expenseService.FetchLists(userID, "", "")
	}

	for _, expense := range fetchedExpenses {
		expenses = append(expenses, api.Expense{
			Amount: expense.Amount,
			Category: int(expense.Category),
			Description: expense.Description,
			Id: strconv.Itoa(int(expense.ID)),
			PaidAt: openapi_types.Date{Time: expense.PaidAt},
		})
	}

	return api.GetExpenses200JSONResponse{FetchExpenseListsResponseJSONResponse: api.FetchExpenseListsResponseJSONResponse{Expenses: expenses}}, nil
}

func (eh *expensesHandler) GetExpensesTotalAmounts(ctx context.Context, request api.GetExpensesTotalAmountsRequestObject) (api.GetExpensesTotalAmountsResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	fromDate := request.Params.FromDate
	toDate := request.Params.ToDate

	totalAmounts := eh.expenseService.FetchTotalAmount(userID, fromDate, toDate)

	var resTotalAmounts []api.TotalAmountLists
	for _, totalAmount := range totalAmounts {
		resTotalAmounts = append(resTotalAmounts, api.TotalAmountLists{
			Date: openapi_types.Date{Time: totalAmount.PaidAt},
			ExtendProps: struct{TotalAmount int "json:\"totalAmount\""; Type string "json:\"type\""} {
				TotalAmount: totalAmount.TotalAmount,
				Type: "expense",
			},
		})
	}
	return api.GetExpensesTotalAmounts200JSONResponse{TotalAmountListsResponseJSONResponse: api.TotalAmountListsResponseJSONResponse{TotalAmounts: resTotalAmounts}}, nil
}

func (eh *expensesHandler) PostExpenses(ctx context.Context, request api.PostExpensesRequestObject) (api.PostExpensesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	createdExpense, validationErr := eh.expenseService.Create(userID, request.Body)

	resValidationError := eh.expenseService.MappingValidationErrorStruct(validationErr)

	return api.PostExpenses200JSONResponse{StoreExpenseResponseJSONResponse: api.StoreExpenseResponseJSONResponse{
		Errors: resValidationError,
		Expense: api.Expense{
			Id: strconv.Itoa(createdExpense.ID),
			PaidAt: openapi_types.Date{Time: createdExpense.PaidAt},
			Amount: createdExpense.Amount,
			Category: int(createdExpense.Category),
			Description: createdExpense.Description,
		},
	}}, nil
}
