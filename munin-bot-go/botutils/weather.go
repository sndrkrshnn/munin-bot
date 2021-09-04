package botutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherContent struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type WeatherMetrics struct {
	Temperature float32 `json:"temp"`
	Feelslike   float32 `json:"feels_like"`
	Humidity    float32 `json:"humidity"`
}
type Wind struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Weather []WeatherContent `json:"weather"`
	Metrics WeatherMetrics   `json:"main"`
	Wind    Wind             `json:"wind"`
}

func GetWeather(city string, API_KEY string) Weather {
	getURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=Vaxjo&units=metric&appid=%s", API_KEY)
	if city != "" {
		getURL = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, API_KEY)
	}

	resp, err := http.Get(getURL)
	if err != nil {
		log.Fatal(err)
	}

	bytes, readErr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readErr != nil {
		log.Fatal(readErr)
	}

	var weather Weather
	errUnmarshal := json.Unmarshal(bytes, &weather)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return weather
}
