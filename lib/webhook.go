package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// WebhookValidate : validate webhook url
func WebhookValidate(webhookURL string) error {
	var data map[string]interface{}

	if strings.Contains(webhookURL, "discordapp.com") {
		data = map[string]interface{}{
			"username":   "편지봇",
			"avatar_url": "https://cdn.discordapp.com/attachments/683175932873539589/689459371151065088/message-3592640_1280.jpg",
			"content":    "🥳 편지봇을 구독해주셔서 감사합니다!\n🗓️ 매일 아침 8시에 이 채널로 브리핑을 해드릴께요.",
		}
	} else if strings.Contains(webhookURL, "hooks.slack.com") {
		data = map[string]interface{}{
			"color":   "#928BFF",
			"pretext": "🥳 편지봇을 구독해주셔서 감사합니!\n🗓️ 매일 아침 8시에 이 채널로 브리핑을 해드릴께요",
		}
	}

	return sendWebhook(webhookURL, data)
}

func sendWebhook(webhookURL string, data interface{}) error {
	jsonByte, _ := json.Marshal(data)
	buff := bytes.NewBuffer(jsonByte)

	resp, err := http.Post(webhookURL, "application/json", buff)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errors.New("Undefined WebhookUrl")
	}

	return err
}
