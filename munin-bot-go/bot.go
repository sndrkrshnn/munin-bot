package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const API_KEY = "560eb9ae5c3d4fcea61d6c85ece0317a"
const TELE_BOT_TOKEN = "1921970606:AAFvOb2DLn58gQqaBGXy2R4a5PFewMcP5NE"

type Content struct {
	Articles []Article `json:"articles"`
}
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func HandleError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func main() {
	bot, error := tgbot.NewBotAPI(TELE_BOT_TOKEN)
	HandleError(error)

	bot.Debug = false

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
			msg.Text = "I am Munin, Odin's raven. I gather news from Midgard when commanded /getnews."
			if _, err := bot.Send(msg); err != nil {
				log.Fatal(err)
			}
		case "getnews":
			for i := 0; i < 5; i++ {
				msg.Text = fmt.Sprintf("Title: %+s\nURL: %+s", GetContent().Articles[i].Title, GetContent().Articles[i].URL)
				if _, err := bot.Send(msg); err != nil {
					log.Fatal(err)
				}
			}
		case "kaw":
			msg.Text = "You expect me to kaw, cus I am a raven? :|"
			if _, err := bot.Send(msg); err != nil {
				log.Fatal(err)
			}
		default:
			msg.Text = "Kaw kaw idk what you kawing?"
			if _, err := bot.Send(msg); err != nil {
				log.Fatal(err)
			}
		}

	}

}

func GetContent() Content {
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=" + API_KEY)

	HandleError(err)

	bytes, readErr := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	HandleError(readErr)

	var c Content
	// Unmarshal is basically unpacking a string to a JSON based struct
	errUnmarshal := json.Unmarshal(bytes, &c)
	HandleError(errUnmarshal)
	return c

}
