package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	p "example.com/main/processor"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, error := tgbot.NewBotAPI(os.Getenv("TELE_BOT_TOKEN"))
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
		case "start":
			msg.Text = "I am Munin, Odin's raven. I gather news from Midgard when commanded." +
				"\nSince I roam Midgard I might as well tell you the weather status over there."
		case "help":
			msg.Text = "\nIf you want to gather news for a custom word, use /getnews <keyword>." +
				"\nIf keyword contains more than one word, use /getnews <1word-2word>." +
				"\nYou can also ask me if you need an umbrella by commanding /weather or /weather <cityname>." +
				"\nYou can also suggest new features that you'd like to see in me with /suggestions <suggestion>."
		case "suggestions":
			var s string
			if update.Message.CommandArguments() != "" {
				s = fmt.Sprintf("Someone suggested: %s", strings.ToLower(update.Message.CommandArguments()))
			}
			msg.Text = "Your query is delivered to the ears of the AllFather.\n"
			if _, err := bot.Send(tgbot.NewMessage(1145663468, s)); err != nil {
				panic(err)
			}
		case "getnews":
			var keyword = ""
			if update.Message.CommandArguments() != "" {
				keyword = strings.ToLower(update.Message.CommandArguments())
			}
			msg.Text = p.ProcessNews(keyword, os.Getenv("NEWS_API_KEY"))
		case "weather":
			var city = "Vaxjo"
			if update.Message.CommandArguments() != "" {
				city = strings.ToLower(update.Message.CommandArguments())
			}
			msg.Text = p.ProcessWeather(city, os.Getenv("WEATHER_API_KEY"))
		case "kaw":
			msg.Text = "You expect me to kaw, cus I am a raven? :|"
		default:
			msg.Text = "Kaw kaw idk what you kawing?"
		}
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
