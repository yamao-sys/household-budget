package handlers

import (
	api "apps/api"
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

type mainHandler struct {
	csrfHandler CsrfHandler
	usersHandler UsersHandler
	expensesHandler ExpensesHandler
}

func NewMainHandler(
	csrfHandler CsrfHandler,
	usersHandler UsersHandler,
	expensesHandler ExpensesHandler,
) MainHandler {
	return &mainHandler{csrfHandler, usersHandler, expensesHandler}
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

func (mh *mainHandler) GetExpenses(ctx context.Context, request api.GetExpensesRequestObject) (api.GetExpensesResponseObject, error) {
	res, err := mh.expensesHandler.GetExpenses(ctx, request)
	return res, err
}

func (mh *mainHandler) GetExpensesTotalAmounts(ctx context.Context, request api.GetExpensesTotalAmountsRequestObject) (api.GetExpensesTotalAmountsResponseObject, error) {
	res, err := mh.expensesHandler.GetExpensesTotalAmounts(ctx, request)
	return res, err
}

func (mh *mainHandler) PostExpenses(ctx context.Context, request api.PostExpensesRequestObject) (api.PostExpensesResponseObject, error) {
	res, err := mh.expensesHandler.PostExpenses(ctx, request)
	return res, err
}
