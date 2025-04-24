package main

import (
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

	// NOTE: Handlerをルーティングに追加
	// e := businessmiddlewares.ApplyMiddlewares(echo.New())
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Household Budget!")
	})

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
