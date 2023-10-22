package main

import (
	"fmt"
	"net/http"
	"os"
	"server/config"
	"server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Hello World")

	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_URL")},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin},
	}))

	db := config.NewDB()

	config.AutoMigrate(db)
	handler.SetupRoutes(e, db)
	// ポート番号を環境変数から読み取り、デフォルトのポート番号を設定
	port := os.Getenv("PORT")
	if port == "" {
			port = "1222" // デフォルトのポート番号
	}

	e.Logger.Fatal(e.Start(":" + port))
	// e.Logger.Fatal(e.Start(":1222"))
}
