package models

type WeatherData struct {
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base       string  `json:"base"`
	Main       Main    `json:"main"`
	Visibility int     `json:"visibility"`
	Wind       Wind    `json:"wind"`
	Clouds     Clouds  `json:"clouds"`
	Dt         int     `json:"dt"`
	Timezone   int     `json:"timezone"`
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Cod        int     `json:"cod"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}