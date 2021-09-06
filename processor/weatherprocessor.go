package processor

import (
	"fmt"

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
	switch t := w.Metrics.Feelslike; {
	case t <= 10.0 && (!(t < 6.0) || !(t < 2.0)):
		temp += "Might wanna consider wearing a sweater or something\n" +
			fmt.Sprintf("The temperature outside is %.1f°C\n", t)
	case t <= 6.0 && !(t < 2.0):
		temp += "Definitely wear a sweater\n" +
			fmt.Sprintf("The temperature outside is %.1f°C\n", t)
	case t <= 2.0:
		temp += "Wear armor if you will.. The only objective is to make yourself warm XD\n" +
			fmt.Sprintf("The temperature outside is %.1f°C\n", t)
	case t > 10.0 && (!(t > 15.0)):
		temp += "You won't need a sweater per se.. but can take it with you if you want\n" +
			fmt.Sprintf("The temperature outside is %.1f°C\n", t)
	case t >= 15.0:
		temp += "Show off that body.. Humans were never meant to wear clothes..\n" +
			fmt.Sprintf("The temperature outside is %.1f°C\n", t)
	default:
		temp += fmt.Sprintf("Hmm.. I suppose Odin hasn't programmed that temperature range yet. So it's just %.1f°C\n", t)
	}
	return temp
}
