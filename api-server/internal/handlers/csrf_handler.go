package handlers

import (
	api "apps/api"
	"context"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
)

type CsrfHandler interface {
	GetCsrf(ctx context.Context, request api.GetCsrfRequestObject) (api.GetCsrfResponseObject, error)
}

type csrfHandler struct {}

func NewCsrfHandler() CsrfHandler {
	return &csrfHandler{}
}

func (ch *csrfHandler) GetCsrf(ctx context.Context, request api.GetCsrfRequestObject) (api.GetCsrfResponseObject, error) {
	csrfToken, ok := ctx.Value(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return api.GetCsrf500JSONResponse{InternalServerErrorResponseJSONResponse: api.InternalServerErrorResponseJSONResponse{
			Code: http.StatusInternalServerError,
			Message: "failed to retrieval token",
		}}, nil
	}
	
	return api.GetCsrf200JSONResponse{CsrfResponseJSONResponse: api.CsrfResponseJSONResponse{ CsrfToken: csrfToken }}, nil
}
