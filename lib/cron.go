package lib

import (
	"time"
)

// Cron : Run Send Letter Cron Job
func Cron() {
	ticker := time.NewTicker(time.Minute)

	go func() {
		for t := range ticker.C {
			if t.Hour() == 8 && t.Minute() == 0 {

			}
		}
	}()
}
