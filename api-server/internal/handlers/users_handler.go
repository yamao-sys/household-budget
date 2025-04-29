package handlers

import (
	"apps/api"
	"apps/internal/helpers"
	"apps/internal/services"
	"context"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UsersHandler interface {
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
}

type usersHandler struct {
	userService services.UserService
}

func NewUsersHandler(userService services.UserService) UsersHandler {
	return &usersHandler{userService}
}

func (uh *usersHandler) PostUsersValidateSignUp(ctx context.Context, request api.PostUsersValidateSignUpRequestObject) (api.PostUsersValidateSignUpResponseObject, error) {
	inputs := api.PostUsersSignUpJSONRequestBody{
		Name:     request.Body.Name,
		Email:    request.Body.Email,
		Password: request.Body.Password,
	} 
	err := uh.userService.ValidateSignUp(ctx, inputs)
	validationError := uh.mappingValidationErrorStruct(err)

	res := &api.UserSignUpResponse{
		Code: http.StatusOK,
		Errors: validationError,
	}
	return api.PostUsersValidateSignUp200JSONResponse{UserSignUpResponseJSONResponse: api.UserSignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (uh *usersHandler) PostUsersSignUp(ctx context.Context, request api.PostUsersSignUpRequestObject) (api.PostUsersSignUpResponseObject, error) {
	err := uh.userService.ValidateSignUp(ctx, *request.Body)
	if err != nil {
		validationError := uh.mappingValidationErrorStruct(err)
	
		res := &api.UserSignUpResponse{
			Code: http.StatusOK,
			Errors: validationError,
		}
		return api.PostUsersSignUp200JSONResponse{UserSignUpResponseJSONResponse: api.UserSignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
	}

	signUpErr := uh.userService.SignUp(ctx, *request.Body)
	if signUpErr != nil {
		return api.PostUsersSignUp500JSONResponse{InternalServerErrorResponseJSONResponse: api.InternalServerErrorResponseJSONResponse{Code: http.StatusInternalServerError, Message: signUpErr.Error()}}, nil
	}

	res := &api.UserSignUpResponse{
		Code: http.StatusOK,
		Errors: api.UserSignUpValidationError{},
	}
	return api.PostUsersSignUp200JSONResponse{UserSignUpResponseJSONResponse: api.UserSignUpResponseJSONResponse{Code: res.Code, Errors: res.Errors}}, nil
}

func (uh *usersHandler) PostUsersSignIn(ctx context.Context, request api.PostUsersSignInRequestObject) (api.PostUsersSignInResponseObject, error) {
	statusCode, tokenString, err := uh.userService.SignIn(ctx, *request.Body)
	switch (statusCode) {
	case http.StatusInternalServerError:
		return api.PostUsersSignIn500JSONResponse{InternalServerErrorResponseJSONResponse: api.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		}}, nil
	case http.StatusBadRequest:
		return api.PostUsersSignIn400JSONResponse{UserSignInBadRequestResponseJSONResponse: api.UserSignInBadRequestResponseJSONResponse{
			Errors: []string{err.Error()},
		}}, nil
	}
	
	// NOTE: Cookieにtokenをセット
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   3600 * 24,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	return api.PostUsersSignIn200JSONResponse{UserSignInOkResponseJSONResponse: api.UserSignInOkResponseJSONResponse{
		Headers: api.UserSignInOkResponseResponseHeaders{
			SetCookie: cookie.String(),
		},
	}}, nil
}

func (uh *usersHandler) GetUsersCheckSignedIn(ctx context.Context, request api.GetUsersCheckSignedInRequestObject) (api.GetUsersCheckSignedInResponseObject, error) {
	userID, _ := helpers.ExtractUserID(ctx)
	return api.GetUsersCheckSignedIn200JSONResponse(uh.userService.ExistsUser(userID)), nil
}

func (uh *usersHandler) mappingValidationErrorStruct(err error) api.UserSignUpValidationError {
	var validationError api.UserSignUpValidationError
	if err == nil {
		return validationError
	}

	if errors, ok := err.(validation.Errors); ok {
		// NOTE: レスポンス用の構造体にマッピング
		for field, err := range errors {
			messages := []string{err.Error()}
			switch field {
			case "name":
				validationError.Name = &messages
			case "email":
				validationError.Email = &messages
			case "password":
				validationError.Password = &messages
			}
		}
	}
	return validationError
}
