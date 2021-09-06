package processor

import (
	"fmt"

	"example.com/main/botutils"
)

func ProcessNews(keyword string, API_KEY string) string {
	c := botutils.GetContent(keyword, API_KEY)
	var reply string
	if len(c.Articles) > 0 {
		for i := 0; i < 3; i++ {
			reply += fmt.Sprintf("Title: %+s\nURL: %+s\n", c.Articles[i].Title, c.Articles[i].URL)
		}
	} else {
		reply += "I've searched far and wide to no avail.. My sources have restricted me for 24 hours as I bothered them too much :<\n"
	}
	return reply
}
