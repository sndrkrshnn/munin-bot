package responses

import (
	"fmt"
	"math/rand"
	"time"
	"example.com/main/types"
)
func HandleWeatherResponse(s string) string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(getFog()) - 1
	random := rand.Intn(max-min+1) + min
	switch s {
	case "Fog", "fog", "Mist", "mist":
		return getFog()[random]
	case "Rain", "rain":
		return getRain()[random]
	case "Clear", "clear":
		return getClear()[random]
	case "Cloudy", "cloudy", "Clouds", "clouds":
		return getCloudy()[random]
	case "Haze", "haze":
		return getHaze()[random]
	default:
		return fmt.Sprintf("Hmm.. I suppose Odin hasn't programmed that part of the weather yet. So it's just %s\n", s)	
	}
}

func HandleTempResponse(t types.Weather) string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(getFog()) - 1
	random := rand.Intn(max-min+1) + min
	switch s := t.Metrics.Feelslike; {
	case s <= 10.0 && (!(s < 6.0) || !(s < 2.0)):
		return getLightCold()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	case s <= 6.0 && !(s < 2.0):
		return getAverageCold()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	case s <= 2.0 && !(s < 0.0):
		return getFreezingCold()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	case s < 0.0:
		return getFreezingCold()[random] + fmt.Sprintf("The temeature outside is -%0.1f°C\n", s)
	case s > 10.0 && (!(s > 15.0)) && !(t.Wind.Speed >= 10.0) :
		return getWarm()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	case s > 10.0 && (!(s > 15.0)) && t.Wind.Speed >= 10.0:
		return getAverageCold()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n, but it's windy..", s)
	case s >= 15.0 && s <= 19.0 && !(t.Wind.Speed > 10.0):
		return getWarm()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	case s >= 15.0 && s <= 19.0 && t.Wind.Speed > 10.0:
		return getLightCold()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n, but it's windy..", s)
	case s >= 19.0:
		return getBurningHot()[random] + fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	default:
		return fmt.Sprintf("The temperature outside is %0.1f°C\n", s)
	}
}
