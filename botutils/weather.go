package botutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/main/types"
)

func GetWeather(city string, API_KEY string) types.Weather {
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

	var weather types.Weather
	errUnmarshal := json.Unmarshal(bytes, &weather)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return weather
}
