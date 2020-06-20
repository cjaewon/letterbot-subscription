package routes

import (
	"letterbot-subscription/lib"
	"strings"

	"github.com/labstack/echo/v4"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	g.POST("/subscribe", subscribe)
}

func subscribe(c echo.Context) error {
	type RequestBody struct {
		WebhookURL string `json:"url" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	if !(strings.Contains(body.WebhookURL, "discordapp.com") || strings.Contains(body.WebhookURL, "hooks.slack.com")) {
		return c.NoContent(403)
	}

	if err := lib.WebhookValidate(body.WebhookURL); err != nil {
		return c.NoContent(403)
	}

	return nil
}
