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

	lib.SendLetter("https://discordapp.com/api/webhooks/689489066068279329/aDKBCdva3iu-6NYX1I2Qrt6Kyo9i6QFdrZdqfaOFctF1JnnGJKFqfoC4K0Pdr79fKDdp")

	panic(1)
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

	e.Logger.Fatal(e.Start("localhost:3000"))
}
