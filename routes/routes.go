package routes

import (
	"strings"

	"letterbot-subscription/database/models"
	"letterbot-subscription/lib"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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
	db, _ := c.Get("db").(*gorm.DB)

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	if !(strings.Contains(body.WebhookURL, "discordapp.com") || strings.Contains(body.WebhookURL, "hooks.slack.com")) {
		return c.NoContent(403)
	}

	// exits check
	if err := db.Where("url = ?", body.WebhookURL).First(&models.Webhook{}).Error; err == nil {
		return c.NoContent(409)
	}

	if err := lib.WebhookValidate(body.WebhookURL); err != nil {
		return c.NoContent(403)
	}

	webhook := models.Webhook{
		URL: body.WebhookURL,
	}

	db.Create(&webhook)
	log.WithField("webhook_url", body.WebhookURL).Info("Create Webhook")

	return nil
}
