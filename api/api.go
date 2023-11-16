package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"github.com/Shimodaira0910/weather/models"
)	

type Weather struct{}

func (w *Weather)GetWeatherInfo(apiEndpoint string, apiKey string, lang string, tmp string, cityName string) (string){
	url := fmt.Sprintf("%s?q=%s&lang=%s&appid=%s&units=%s", apiEndpoint, cityName, lang, apiKey, tmp)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("リクエストエラー:", err)
		return "request error"
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
    if err != nil {
		fmt.Println(err)
        return "read error"
    }
	
	weather := models.WeatherData{}
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Println("存在しない都市名です", err)
		return "not exist city"
	}

	fmt.Println("現在の時刻は" + convertNowtmToString(weather.Timezone))
	fmt.Println(weather.Name + "の天気は" + weather.Weather[0].Description + "です。")
	fmt.Println(weather.Name + "の最高気温は" + strconv.FormatFloat(weather.Main.TempMax, 'f', 2, 64) + "です。")
	return "ok"
}

func convertNowtmToString(unixtime int) string{
	nowTime := time.Now()
	toStringNowTime := nowTime.Format("15:04:05")	
	return toStringNowTime
}

