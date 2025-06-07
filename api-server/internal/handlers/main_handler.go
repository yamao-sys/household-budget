package handlers

import (
	api "apps/apis"
	"context"
)

type MainHandler interface {
	// handlers /csrf
	GetCsrf(ctx context.Context, request api.GetCsrfRequestObject) (api.GetCsrfResponseObject, error)

	// User SignIn
	// (POST /users/signIn)
	PostUsersSignIn(ctx context.Context, request api.PostUsersSignInRequestObject) (api.PostUsersSignInResponseObject, error)
	// User SignUp
	// (POST /users/signUp)
	PostUsersSignUp(ctx context.Context, request api.PostUsersSignUpRequestObject) (api.PostUsersSignUpResponseObject, error)
	// User Validate SignUp
	// (POST /users/validateSignUp)
	PostUsersValidateSignUp(ctx context.Context, request api.PostUsersValidateSignUpRequestObject) (api.PostUsersValidateSignUpResponseObject, error)
	// User CheckSignedIn
	// (GET /users/checkSignedIn)
	GetUsersCheckSignedIn(ctx context.Context, request api.GetUsersCheckSignedInRequestObject) (api.GetUsersCheckSignedInResponseObject, error)

	// Get Expenses
	// (GET /expenses)
	GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error)
	// Get Expenses TotalAmounts
	// (GET /expenses/totalAmounts)
	GetExpensesTotalAmounts(ctx context.Context, request api.GetExpensesTotalAmountsRequestObject) (api.GetExpensesTotalAmountsResponseObject, error)
	// Get Expenses CategoryTotalAmounts
	// (GET /expenses/categoryTotalAmounts)
	GetExpensesCategoryTotalAmounts(ctx context.Context, request api.GetExpensesCategoryTotalAmountsRequestObject) (api.GetExpensesCategoryTotalAmountsResponseObject, error)
	// Post Expense
	// (POST /expenses)
	PostExpenses(ctx context.Context, request api.PostExpensesRequestObject) (api.PostExpensesResponseObject, error)

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

type mainHandler struct {
	csrfHandler CsrfHandler
	usersHandler UsersHandler
	expensesHandler ExpensesHandler
	incomesHandler IncomesHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	usersHandler UsersHandler,
	expensesHandler ExpensesHandler,
	incomesHandler IncomesHandler,
) MainHandler {
	return &mainHandler{csrfHandler, usersHandler, expensesHandler, incomesHandler}
}

func (mh *mainHandler) GetCsrf(ctx context.Context, request api.GetCsrfRequestObject) (api.GetCsrfResponseObject, error) {
	res, err := mh.csrfHandler.GetCsrf(ctx, request)
	return res, err
}

func (mh *mainHandler) PostUsersSignIn(ctx context.Context, request api.PostUsersSignInRequestObject) (api.PostUsersSignInResponseObject, error) {
	res, err := mh.usersHandler.PostUsersSignIn(ctx, request)
	return res, err
}

func (mh *mainHandler) PostUsersSignUp(ctx context.Context, request api.PostUsersSignUpRequestObject) (api.PostUsersSignUpResponseObject, error) {
	res, err := mh.usersHandler.PostUsersSignUp(ctx, request)
	return res, err
}

func (mh *mainHandler) PostUsersValidateSignUp(ctx context.Context, request api.PostUsersValidateSignUpRequestObject) (api.PostUsersValidateSignUpResponseObject, error) {
	res, err := mh.usersHandler.PostUsersValidateSignUp(ctx, request)
	return res, err
}

func (mh *mainHandler) GetUsersCheckSignedIn(ctx context.Context, request api.GetUsersCheckSignedInRequestObject) (api.GetUsersCheckSignedInResponseObject, error) {
	res, err := mh.usersHandler.GetUsersCheckSignedIn(ctx, request)
	return res, err
}

func (mh *mainHandler) GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error) {
	res, err := mh.expensesHandler.GetExpenses(ctx, request)
	return res, err
}

func (mh *mainHandler) GetExpensesTotalAmounts(ctx context.Context, request api.GetExpensesTotalAmountsRequestObject) (api.GetExpensesTotalAmountsResponseObject, error) {
	res, err := mh.expensesHandler.GetExpensesTotalAmounts(ctx, request)
	return res, err
}

func (mh *mainHandler) GetExpensesCategoryTotalAmounts(ctx context.Context, request api.GetExpensesCategoryTotalAmountsRequestObject) (api.GetExpensesCategoryTotalAmountsResponseObject, error) {
	res, err := mh.expensesHandler.GetExpensesCategoryTotalAmounts(ctx, request)
	return res, err
}

func (mh *mainHandler) PostExpenses(ctx context.Context, request api.PostExpensesRequestObject) (api.PostExpensesResponseObject, error) {
	res, err := mh.expensesHandler.PostExpenses(ctx, request)
	return res, err
}

func (mh *mainHandler) GetIncomes(ctx context.Context, request api.GetIncomesRequestObject) (api.GetIncomesResponseObject, error) {
	res, err := mh.incomesHandler.GetIncomes(ctx, request)
	return res, err
}

func (mh *mainHandler) GetIncomesTotalAmounts(ctx context.Context, request api.GetIncomesTotalAmountsRequestObject) (api.GetIncomesTotalAmountsResponseObject, error) {
	res, err := mh.incomesHandler.GetIncomesTotalAmounts(ctx, request)
	return res, err
}

func (mh *mainHandler) GetIncomesClientTotalAmounts(ctx context.Context, request api.GetIncomesClientTotalAmountsRequestObject) (api.GetIncomesClientTotalAmountsResponseObject, error) {
	res, err := mh.incomesHandler.GetIncomesClientTotalAmounts(ctx, request)
	return res, err
}

func (mh *mainHandler) PostIncomes(ctx context.Context, request api.PostIncomesRequestObject) (api.PostIncomesResponseObject, error) {
	res, err := mh.incomesHandler.PostIncomes(ctx, request)
	return res, err
}
