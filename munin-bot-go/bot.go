package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const API_KEY = "560eb9ae5c3d4fcea61d6c85ece0317a"
const TELE_BOT_TOKEN = "1921970606:AAFvOb2DLn58gQqaBGXy2R4a5PFewMcP5NE"

type Content struct {
	Articles []Article `json:"articles"`
}
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

func GetContent(keyword string) Content {
	getURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=%s", API_KEY)
	if keyword != "" {
		getURL = fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", keyword, API_KEY)
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

	var c Content
	errUnmarshal := json.Unmarshal(bytes, &c)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
	return c
}

func main() {
	bot, error := tgbot.NewBotAPI(TELE_BOT_TOKEN)
	if error != nil {
		log.Fatal(error)
	}
	updateConfig := tgbot.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}
		msg := tgbot.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help":
			msg.Text = "I am Munin, Odin's raven. I gather news from Midgard when commanded /getnews." +
				"\nIf you want to search for a custom word, use /getnews <keyword>." +
				"\nIf keyword contains more than one word, use /getnews <1word-2word>."
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		case "getnews":
			var keyword = ""
			if update.Message.CommandArguments() != "" {
				keyword = strings.ToLower(update.Message.CommandArguments())
			}
			for i := 0; i < 5; i++ {
				msg.Text = fmt.Sprintf("Title: %+s\nURL: %+s", GetContent(keyword).Articles[i].Title, GetContent(keyword).Articles[i].URL)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			}
		case "kaw":
			msg.Text = "You expect me to kaw, cus I am a raven? :|"
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		default:
			msg.Text = "Kaw kaw idk what you kawing?"
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

	}

}
