package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"github.com/Shimodaira0910/weather/models"
)	

type Weather struct{}

// api呼び出し
func (w *Weather)GetWeatherInfo(apiEndpoint string, apiKey string, lang string, tmp string, cityName string) (string, error){
	url := fmt.Sprintf("%s?q=%s&lang=%s&appid=%s&units=%s", apiEndpoint, cityName, lang, apiKey, tmp)

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request error: %v", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
    if err != nil {
		return "", fmt.Errorf("read error: %v", err)
    }
	
	weather := models.WeatherData{}
	if err := json.Unmarshal(body, &weather); err != nil {
		return "", fmt.Errorf("invalid city name: %v", err)
	}

	printWeatherInfo(weather)
	return "", nil
}

// 現在時刻の取得
func convertNowtmToString(unixtime int) string{
	nowTime := time.Now()
	toStringNowTime := nowTime.Format("15:04:05")	
	return toStringNowTime
}

// 取得した天気のデータを出力する
func printWeatherInfo(weather models.WeatherData){
	fmt.Printf("現在の時刻は%sです。\n" , convertNowtmToString(weather.Timezone))
	fmt.Printf("%sの天気は%sです。\n",weather.Name,weather.Weather[0].Description)
	fmt.Printf("%sの最高気温は%.2fです。\n", weather.Name, weather.Main.TempMax)
}