package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	p "example.com/main/processor"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

const NEWS_API_KEY = "560eb9ae5c3d4fcea61d6c85ece0317a"
const TELE_BOT_TOKEN = "1921970606:AAFvOb2DLn58gQqaBGXy2R4a5PFewMcP5NE"
const WEATHER_API_KEY = "b2b1c07c6349055ee36c756e00b7ca4c"

func main() {
	bot, error := tgbot.NewBotAPI(TELE_BOT_TOKEN)
	if error != nil {
		log.Fatal(error)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set webhook url (curl)
	log.Print("Listening on: " + os.Getenv("PORT"))

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	for update := range updates {
		if update.Message == nil {
			return
		}
		log.Print(update.Message.Chat.ID)
		msg := tgbot.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "help", "intro":
			log.Print("gotcha home!!")
			msg.Text = "I am Munin, Odin's raven. I gather news from Midgard when commanded /getnews." +
				"\nIf you want to search for a custom word, use /getnews <keyword>." +
				"\nIf keyword contains more than one word, use /getnews <1word-2word>." +
				"\nYou can also ask me if you need a raincoat by commanding /dinar or /dinar <cityname>."
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		case "getnews":
			var keyword = ""
			if update.Message.CommandArguments() != "" {
				keyword = strings.ToLower(update.Message.CommandArguments())
			}
			msg.Text = p.ProcessNews(keyword, NEWS_API_KEY)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		case "dinar":
			var city = "Vaxjo"
			if update.Message.CommandArguments() != "" {
				city = strings.ToLower(update.Message.CommandArguments())
			}
			msg.Text = p.ProcessWeather(city, WEATHER_API_KEY)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
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
