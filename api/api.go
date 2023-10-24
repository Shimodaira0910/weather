package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Shimodaira0910/weather/env"
	"github.com/Shimodaira0910/weather/models"
)	

type Weather struct{}

func (w *Weather)GetWeatherInfo(cityName string) (string, error){
	env := env.Env{}
	apiEndpoint := env.LoadEnv("API_ENDPOINT")
	apiKey := env.LoadEnv("WEATHER_API_TOKEN") 

	url := fmt.Sprintf("%s?q=%s&lang=%s&appid=%s", apiEndpoint, cityName, "ja", apiKey)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("リクエストエラー:", err)
		return "", err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
    if err != nil {
        return "", err
    }
	
	weather := models.WeatherData{}
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Printf("存在しない都市名です")
		return "", err
	}

	fmt.Println("現在の時刻は" + convertUnixtmToString(weather.Timezone))
	fmt.Println(weather.Name + "の天気は" + weather.Weather[0].Description + "です。")
	return "", err
}

func convertUnixtmToString(unixtime int) string{
	dtFormUnix := time.Unix(int64(unixtime), 0)
	formattedTime := dtFormUnix.Format("15:04:05")
	return formattedTime
}