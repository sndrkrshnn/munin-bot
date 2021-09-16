package processor

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/main/botutils"
)

func ProcessImage(keyword string, API_KEY string) string {
	p := botutils.GetImageContent(keyword, API_KEY)
	var reply string
	if p.Photos.Total > 0 {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(p.Photos.Photo) - 1
		random := rand.Intn(max-min+1) + min
		reply += fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_c.jpg", p.Photos.Photo[random].Server, p.Photos.Photo[random].ID, p.Photos.Photo[random].Secret)
	} else {
		reply += "I've searched far and wide to no avail.. My sources have restricted me for 24 hours as I bothered them too much :<\n"
	}
	return reply
}
