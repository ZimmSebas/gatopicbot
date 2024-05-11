package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Response []struct {
	Breeds []any  `json:"breeds"`
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func cat_request(CAT_API_TOKEN string) string {
	client := &http.Client{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	request, err := http.NewRequest("GET", "https://api.thecatapi.com/v1/images/search?format=json&limit=1", nil)

	if err != nil {
		log.Fatalf("Error creating the request %v", err)
	}

	// TO-DO: Set Cat API TOKEN
	request.Header.Set("x-api-key", CAT_API_TOKEN)

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Request failed with error %v", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObj Response
	json.Unmarshal(responseData, &responseObj)

	return responseObj[0].Url
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//Env tokens
	CAT_API_TOKEN := os.Getenv("CAT_API_TOKEN")
	TELEGRAM_BOT_API_TOKEN := os.Getenv("TELEGRAM_BOT_API_TOKEN")

	//Basic messages
	start_message := "Hola! Este es un bot de gatitos, para recibir un michi 🐱, escribí /michi"
	default_message := "Perdón, no entendí tu mensaje, si querés recibir un michi 🐱, escribí /michi"
	var msg_text string

	bot, err := tgbotapi.NewBotAPI(TELEGRAM_BOT_API_TOKEN)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				msg_text = start_message
			case "/michi":
				msg_text = cat_request(CAT_API_TOKEN)
			default:
				msg_text = default_message
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msg_text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
