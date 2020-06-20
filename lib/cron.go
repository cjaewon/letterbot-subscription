package lib

import (
	"letterbot-subscription/database/models"
	"time"

	"github.com/jinzhu/gorm"
)

// Cron : Run Send Letter Cron Job
func Cron(db *gorm.DB) {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for t := range ticker.C {
			if t.Hour() == 8 && t.Minute() == 0 {
				var webhooks []models.Webhook
				db.Find(&webhooks)

				for _, webhook := range webhooks {
					go SendLetter(webhook.URL)
				}
			}
		}
	}()
}
