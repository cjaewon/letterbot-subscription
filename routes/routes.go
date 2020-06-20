package routes

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	g.GET("/subscribe", subscribe)
}

func subscribe(c echo.Context) error {
	type RequestBody struct {
		URL string `json:"url" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	if !strings.Contains(body.URL, "discordapp.com") || !strings.Contains(body.URL, "hooks.slack.com") {
		return c.NoContent(403)
	}

	return nil
}
