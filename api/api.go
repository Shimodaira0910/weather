package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"

	"github.com/Shimodaira0910/weather/env"
)	

type Weather struct{
	Main string `json:"main"`
}

type ApiResponse struct {
	Weather []Weather `json:"weather"`
}

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
	
	var apiRes ApiResponse 
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return "", err
	}

	fmt.Println(apiRes.Weather[0].Main)
	return "", err
}

