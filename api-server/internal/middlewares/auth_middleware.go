package middlewares

import (
	"apps/api"
	"apps/internal/helpers"
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(f api.StrictHandlerFunc, operationID string) api.StrictHandlerFunc {
    return func(ctx echo.Context, i interface{}) (interface{}, error) {
		if !needsAuthenticate(operationID) {
			// NOTE: 認証が不要なURIは認証をスキップ
			return f(ctx, i)
		}

        // NOTE: Cookieからtokenを取得し、JWTの復号
		tokenString, _ := ctx.Cookie("token")
		if tokenString == nil {
			return nil, echo.ErrUnauthorized
		}

		token, _ := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_TOKEN_KEY")), nil
		})

		// NOTE: ログイン種別に応じ、IDをContextにセットする
		c, err := newWithAuthenticateContext(token, ctx)
		if err != nil {
			return nil, echo.ErrUnauthorized
		}

		ctx.SetRequest(ctx.Request().WithContext(c))
        return f(ctx, i)
    }
}

func needsAuthenticate(operationID string) (bool) {
	spec, _ := api.GetSwagger()
	for _, pathItem := range spec.Paths.Map() {
		for _, op := range pathItem.Operations() {
			if op.OperationID != operationID {
				continue
			}
			return len(*op.Security) > 0
		}
	}
	return false
}

func newWithAuthenticateContext(token *jwt.Token, ctx echo.Context) (context.Context, error) {
	var authenticateID int
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authenticateID = int(claims["user_id"].(float64))
	}

	c := helpers.NewWithUserIDContext(ctx.Request().Context(), authenticateID)
	return c, nil
}
