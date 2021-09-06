package processor

import (
	"example.com/main/botutils"
	"example.com/main/responses"
	"example.com/main/types"
)

func ProcessWeather(city string, API_KEY string) string {
	w := botutils.GetWeather(city, API_KEY)
	reply := processMainForecast(w)
	reply += processTemperature(w)
	return reply
}

func processMainForecast(w types.Weather) string {
	var forecast string
	for i := range w.Weather {
		forecast += responses.HandleWeatherResponse(w.Weather[i].Main)
	}
	return forecast
}

func processTemperature(w types.Weather) string {
	var temp string
	temp += responses.HandleTempResponse(w.Metrics.Feelslike)
	return temp
}
