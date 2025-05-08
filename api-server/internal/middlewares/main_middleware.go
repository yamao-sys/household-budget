package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyMiddlewares(e *echo.Echo) *echo.Echo {
	// NOTE: CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_ORIGIN")},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken, echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowOrigin},
		AllowCredentials: true,
	}))

	// NOTE: CSRF対策
	csrfConfig := middleware.CSRFConfig{
		TokenLookup: "header:"+echo.HeaderXCSRFToken,
		CookieMaxAge: 3600,
		CookieSameSite: http.SameSiteNoneMode,
		CookieHTTPOnly: true,
		CookiePath: "/",
		CookieSecure: os.Getenv("APP_ENV") == "production",
		ErrorHandler: func(err error, c echo.Context) error {
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		},
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))

	// NOTE: CSRFトークンをcontext.Contextに埋め込むミドルウェアを適用
	// 	   : StrictHandlerだとecho.Contextがhandler側で使えずのため
	e.Use(csrfContextMiddleware)

	// NOTE: Panicが発生してもサーバを停止することを防ぐ
	e.Use(middleware.Recover())

	return e
}

// CSRFContextMiddleware ... CSRFトークンを context.Context に埋め込むミドルウェア
func csrfContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// NOTE: EchoのcontextからCSRFトークンを取得
		token, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve CSRF token")
		}

		// NOTE: context.Context に CSRF トークンを埋め込む
		//lint:ignore SA1029 It's ok because ContextKey
		ctx := context.WithValue(c.Request().Context(), middleware.DefaultCSRFConfig.ContextKey, token)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
