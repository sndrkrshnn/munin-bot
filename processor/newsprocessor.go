package processor

import (
	"fmt"

	"example.com/main/botutils"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ProcessNews(keyword string, API_KEY string, bot *tgbot.BotAPI, msg tgbot.MessageConfig) {
	c := botutils.GetContent(keyword, API_KEY)
	if len(c.Articles) > 0 {
		for i := 0; i < 3; i++ {
			msg.Text += fmt.Sprintf("Title: %+s\nURL: %+s\n", c.Articles[i].Title, c.Articles[i].URL)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	} else {
		msg.Text += "I've searched far and wide to no avail.. My sources have restricted me for 24 hours as I bothered them too much :<\n"
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
