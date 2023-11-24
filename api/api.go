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

// api呼び出し 天気、最高気温、最低気温
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
	

	weatherInfo := formatWeatherInfo(weather)
	fmt.Print(weatherInfo)
	return "", nil
}

// その国の時刻を取得
func convertUnixtmToString(unixtime int) string{
	dtFormUnix := time.Unix(int64(unixtime), 0)
	formattedTime := dtFormUnix.Format("15:04:05")
	return formattedTime
}

// 現在時刻の取得
func convertNowtmToString() string{
	nowTime := time.Now()
	toStringNowTime := nowTime.Format("15:04:05")	
	return toStringNowTime
}

// 取得した天気のデータをフォーマット化(このメソッドの責務をフォーマット化のみにするため)
func formatWeatherInfo(weather models.WeatherData) string{
	currentTime := convertUnixtmToString(weather.Timezone)
	weatherDescription := fmt.Sprintf("%sの天気は%sです。\n", weather.Name, weather.Weather[0].Description)
	maxTemperatureInfo := fmt.Sprintf("%sの最高気温は%.2fです。\n", weather.Name, weather.Main.TempMax)
    minTemperatureInfo := fmt.Sprintf("%sの最低気温は%.2fです。\n",weather.Name,weather.Main.TempMin)
	return fmt.Sprintf("現在の時刻は%sです。\n%s%s%s", currentTime, weatherDescription, maxTemperatureInfo, minTemperatureInfo)
}