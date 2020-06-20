package main

import (
	"letterbot-subscription/lib"
	"letterbot-subscription/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "dist")

	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	routes.Routes(e.Group("/api"))

	e.Logger.Fatal(e.Start(":3000"))
}
