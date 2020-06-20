package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// WebhookValidate : validate webhook url
func WebhookValidate(webhookURL string) {
	if strings.Contains(webhookURL, "discordapp.com") {
		sendWebhook(webhookURL, map[string]interface{}{
			"username":   "편지봇",
			"avatar_url": "https://cdn.discordapp.com/attachments/683175932873539589/689459371151065088/message-3592640_1280.jpg",

			"content": `✅ 웹훅 주소 테스트로 발송되는 메세지입니다.`,
		})
	} else if strings.Contains(webhookURL, "hooks.slack.com") {
		sendWebhook(webhookURL, map[string]interface{}{
			"color":   "#928BFF",
			"pretext": "✅ 웹훅 주소 테스트로 발송되는 메세지입니다.",
		})
	}
}

func sendWebhook(webhookURL string, data interface{}) error {
	jsonByte, _ := json.Marshal(data)
	buff := bytes.NewBuffer(jsonByte)

	resp, err := http.Post(webhookURL, "application/json", buff)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return err
}
