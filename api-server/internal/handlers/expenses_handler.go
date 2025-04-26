package handlers

import (
	"apps/api"
	"apps/internal/helpers"
	"apps/internal/services"
	"context"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type ExpensesHandler interface {
	// Get Expenses
	// (GET /expenses)
	GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error)
}

type expensesHandler struct {
	expenseService services.ExpenseService
}

func NewExpensesHandler(expenseService services.ExpenseService) ExpensesHandler {
	return &expensesHandler{expenseService}
}

func (eh *expensesHandler) GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	beginningOfMonth := request.Params.BeginningOfMonth

	expenses := eh.expenseService.FetchLists(userID, beginningOfMonth)

	var resExpenses []api.MonthlyCalenderExpense
	for _, expense := range expenses {
		resExpenses = append(resExpenses, api.MonthlyCalenderExpense{
			Date: openapi_types.Date{Time: expense.PaidAt},
			ExtendProps: struct{Amount int "json:\"amount\""; Type string "json:\"type\""}{
				Type: "expense",
				Amount: expense.Amount,
			},
		})
	}
	return api.GetExpenses200JSONResponse{MonthlyCalenderExpenseResponseJSONResponse: api.MonthlyCalenderExpenseResponseJSONResponse{Expenses: &resExpenses}}, nil
}
