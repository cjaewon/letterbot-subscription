package main

import (
	"letterbot-subscription/database"
	"letterbot-subscription/lib"
	"letterbot-subscription/lib/middlewares"
	"letterbot-subscription/routes"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	godotenv.Load(".env")
	lib.LoggerSetting()
}

func main() {
	db := database.Connect()
	defer db.Close()

	lib.Cron(db)

	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}

	e.Static("/", "dist")
	e.Use(middlewares.ContextDB(db))

	routes.Routes(e.Group("/api"))

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
