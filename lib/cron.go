package lib

import (
	"letterbot-subscription/database/models"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Cron : Run Send Letter Cron Job
func Cron(db *gorm.DB) {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for t := range ticker.C {
			if t.Hour() == 8 && t.Minute() == 0 {
				var webhooks []models.Webhook
				db.Find(&webhooks)

				log.WithField("webhook-count", len(webhooks)).Info("Send Webhook Start")
				for _, webhook := range webhooks {
					go SendLetter(webhook.URL)
				}
			}
		}
	}()
}
