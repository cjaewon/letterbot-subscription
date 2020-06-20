package lib

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// GetDate : get date
func GetDate() string {
	now := time.Now()

	year := now.Year()
	month := now.Month()
	day := now.Day()

	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

// GetNews : get news
func GetNews() (discord string, slack string) {
	resp, err := http.Get("https://news.google.com/rss?hl=ko&gl=KR&ceid=KR:ko")
	if err != nil {
		// log
	}

	defer resp.Body.Close()

	type ReponseType struct {
		Channel struct {
			Item []struct {
				Text  string `xml:",chardata"`
				Title string `xml:"title"`
				Link  string `xml:"link"`
			} `xml:"item"`
		} `xml:"channel"`
	}

	var data ReponseType
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bodyBytes, &data)

	for i := 0; i < 3; i++ {
		discord += fmt.Sprintf("[%s](%s)\n", data.Channel.Item[i].Title, data.Channel.Item[i].Link)
		slack += fmt.Sprintf("<%s|%s>\n", data.Channel.Item[i].Title, data.Channel.Item[i].Link)
	}

	return
}

// GetWeather : get weather
func GetWeather() (weather string, temp string) {
	token := os.Getenv("WEATHER_API_KEY")
	city := "Seoul"

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, token)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		// log
	}

	defer resp.Body.Close()

	type ResponseType struct {
		Weather map[int]struct {
			ID int `json:"id"`
		} `json:"weather"`
		Main struct {
			TempMin int `json:"temp_min"`
			TempMax int `json:"temp_max"`
		} `json:"main"`
	}

	var data ResponseType
	json.NewDecoder(resp.Body).Decode(&data)

	weather = weatherMap[data.Weather[0].ID]
	temp = fmt.Sprintf("(%d도 ~ %d도)", data.Main.TempMin, data.Main.TempMax)

	return
}

var weatherMap = map[int]string{
	200: "🌧 가벼운 비를 동반한 천둥구름",
	201: "🌧 비를 동반한 천둥구름",
	202: "🌩 폭우를 동반한 천둥구름",
	210: "🌩 약한 천둥구름",
	211: "🌩 천둥구름",
	212: "🌩 강한 천둥구름",
	221: "🌩 불규칙적 천둥구름",
	230: "🌩 약한 연무를 동반한 천둥구름",
	231: "🌩 연무를 동반한 천둥구름",
	232: "🌧 강한 안개비를 동반한 천둥구름",
	300: "🌧 가벼운 안개비가 내려요",
	301: "🌧 안개비가 내려요",
	302: "🌧 강한 안개비가 내려요",
	310: "🌧 가벼운 적은비가 내려요",
	311: "🌧 적은비가 내려요",
	312: "🌧 강한 적은비가 내려요",
	313: "🌧 소나기와 안개비",
	314: "🌧 강한 소나기와 안개비",
	321: "🌧 소나기가 내려요",
	500: "🌧 약한 비가 내려요",
	501: "🌧 중간 비가 내려요",
	502: "🌧 강한 비가 내려요",
	503: "🌧 매우 강한 비가 내려요",
	504: "🌧 극심한 비가 내려요",
	511: "🌧 우박이 떨어져요",
	520: "🌧 약한 소나기 비가 내려요",
	521: "🌧 소나기 비가 내려요",
	522: "🌧 강한 소나기 비가 내려요",
	531: "🌧 불규칙적 소나기 비가 내려요",
	600: "❄ 가벼운 눈이 내려요",
	601: "❄ 눈이 내려요",
	602: "❄ 강한 눈이 내려요",
	611: "🌧 진눈깨비가 내려요",
	612: "🌧 소나기 진눈깨비가 내려요",
	615: "🌧 약한 비와 눈이 내려요",
	616: "🌧 비와 눈이 내려요",
	620: "🌧 약한 소나기 눈이 내려요",
	621: "🌧 소나기 눈이 내려요",
	622: "❄ 강한 소나기 눈이 내려요",
	701: "박무",
	711: "연기가 있어요",
	721: "⛅ 연무",
	731: "모래 먼지가 날려요",
	741: "안개가 있어요",
	751: "모래가 날려요",
	761: "먼지가 있어요",
	762: "화산재 날려요",
	771: "돌풍이 있어요",
	781: "토네이도",
	800: "☀ 구름 한 점 없는 맑은 하늘입니다.",
	801: "☁ 약간의 구름이 낀 하늘입니다.",
	802: "☁ 드문드문 구름이 낀 하늘입니다.",
	803: "☀ 구름이 거의 없는 하늘입니다.",
	804: "☁ 구름으로 뒤덮인 흐린 하늘입니다.",
	900: "토네이도",
	901: "태풍",
	902: "허리케인",
	903: "한랭",
	904: "♨ 고온",
	905: "💨 바람이 있어요",
	906: "우박이 떨어져요",
	951: "💨 바람이 거의 없어요",
	952: "💨 약한 바람이 있어요",
	953: "💨 부드러운 바람이 있어요",
	954: "💨 중간 세기 바람이 있어요",
	955: "💨 신선한 바람이 있어요",
	956: "💨 센 바람이 있어요",
	957: "💨 돌풍에 가까운 센 바람이 있어요",
	958: "💨 돌풍이 있어요",
	959: "💨 심각한 돌풍이 있어요",
	960: "🌪 폭풍이 발생했어요.",
	961: "🌪 강한 폭풍이 발생했어요.",
	962: "🌪 허리케인",
}
