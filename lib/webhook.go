package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

// SendLetter : send letter
func SendLetter(webhookURL string) {
	date := GetDate()
	discordNews, slackNews := GetNews()
	weather, temp := GetWeather()

	fmt.Println(weather)

	if strings.Contains(webhookURL, "discordapp.com") {
		sendWebhook(webhookURL, map[string]interface{}{
			"username":   "편지봇",
			"avatar_url": "https://cdn.discordapp.com/attachments/683175932873539589/689459371151065088/message-3592640_1280.jpg",
			"content":    fmt.Sprintf("📨 %s 편지가 왔어요!", date),

			"embeds": []map[string]interface{}{
				{
					"fields": []map[string]interface{}{
						{
							"name":   "📅 날짜 / 한국",
							"value":  date,
							"inline": true,
						},
						{
							"name":   "🏞️ 날씨 / 부산",
							"value":  weather,
							"inline": true,
						},
						{
							"name":   "🌡 온도 / 부산",
							"value":  temp,
							"inline": true,
						},
					},
					"footer": map[string]string{
						"text":     "제작자 : 재웜",
						"icon_url": "https://images-ext-2.discordapp.net/external/GyQicPLz_zQO15bOMtiGTtC4Kud7JjQbs1Ecuz7RrtU/https/cdn.discordapp.com/embed/avatars/1.png",
					},
				},
				{
					"title":       "📰 뉴스 / 구글",
					"description": discordNews,
				},
			},
		})
	} else if strings.Contains(webhookURL, "hooks.slack.com") {
		sendWebhook(webhookURL, map[string]interface{}{
			"attachments": []map[string]interface{}{
				{
					"color":   "#928BFF",
					"pretext": fmt.Sprintf("📨 %s 편지가 왔어요!", date),

					"fields": []map[string]interface{}{
						{
							"title": "📅 날짜 / 한국",
							"value": date,
							"short": true,
						},
						{
							"name":  "🏞️ 날씨 / 부산",
							"value": weather,
							"short": true,
						},
						{
							"name":  "🌡 온도 / 부산",
							"value": temp,
							"short": true,
						},
					},
					"footer":      "제작: 재웜",
					"footer_icon": "https://images-ext-2.discordapp.net/external/GyQicPLz_zQO15bOMtiGTtC4Kud7JjQbs1Ecuz7RrtU/https/cdn.discordapp.com/embed/avatars/1.png",
				},
				{
					"fields": []map[string]interface{}{
						{
							"type":  "mrkdwn",
							"title": "📰 뉴스 / 구글",
							"value": slackNews,
						},
					},
				},
			},
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
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errors.New("Undefined WebhookUrl")
	}

	return err
}
