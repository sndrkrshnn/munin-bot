package processor

import (
	"fmt"

	"example.com/main/botutils"
)

func ProcessNews(keyword string, API_KEY string) string {
	c := botutils.GetContent(keyword, API_KEY)
	var reply string
	for i := 0; i < 3; i++ {
		reply += fmt.Sprintf("Title: %+s\nURL: %+s\n", c.Articles[i].Title, c.Articles[i].URL)
	}
	return reply
}
