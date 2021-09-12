package processor

import (
	"fmt"

	"example.com/main/botutils"
)

func ProcessImage(keyword string, API_KEY string) string {
	p := botutils.GetImageContent(keyword, API_KEY)
	var reply string
	if p.ImageContent.Photos.Total > 0 {
		reply += fmt.Sprintf("https://live.staticflickr.com/%s/%s_%s_c.jpg", p.ImageContent.Photos.Photo[0].Server, p.ImageContent.Photos.Photo[0].ID, p.ImageContent.Photos.Photo[0].Secret)
	} else {
		reply += "I've searched far and wide to no avail.. My sources have restricted me for 24 hours as I bothered them too much :<\n"
	}
	return reply
}
