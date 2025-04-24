package main

import (
	"apps/api"
	"apps/database"
	"apps/internal/handlers"
	"apps/internal/middlewares"
	"apps/internal/services"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// NOTE: デプロイ先の環境はSecret Managerで環境変数を管理する
	if os.Getenv("APP_ENV") != "production" {
		loadEnv()
	}

	dbCon := database.Init()

	// NOTE: service層のインスタンス
	userService := services.NewUserService(dbCon)

	// NOTE: Handlerのインスタンス
	csrfHandler := handlers.NewCsrfHandler()
	usersHandler := handlers.NewUsersHandler(userService)

	// NOTE: Handlerをルーティングに追加
	e := middlewares.ApplyMiddlewares(echo.New())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Household Budget!")
	})
	mainHandler := handlers.NewMainHandler(csrfHandler, usersHandler)
	mainStrictHandler := api.NewStrictHandler(mainHandler, []api.StrictMiddlewareFunc{middlewares.AuthMiddleware})
	api.RegisterHandlers(e, mainStrictHandler)

	if err := e.Start(":" + os.Getenv("SERVER_PORT")); err != nil && err != http.ErrServerClosed {
		e.Logger.Errorf("Echo server error: %v", err)
	}
}

func loadEnv() {
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if envFilePath == "" {
		envFilePath = ".env"
	}
	godotenv.Load(envFilePath)
}
