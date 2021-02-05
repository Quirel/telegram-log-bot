package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// init - loads .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

func main() {
	tgToken := os.Getenv("TELEGRAM_LOG_BOT_TOKEN")
	tgBot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatalf("error creating telegram connection: %v\n", err)
	}
	tgLog("Hello", tgBot)
}

// tgLog - logs message to telegram chat
func tgLog(msg string, tgBot *tgbotapi.BotAPI) {
	tgChatId := os.Getenv("TELEGRAM_LOG_BOT_CHAT_ID")
	tgChatIdInt, _ := strconv.ParseInt(tgChatId, 10, 64)
	tgMsg := tgbotapi.NewMessage(tgChatIdInt, "Wa-Go-Bot: "+msg)
	_, err := tgBot.Send(tgMsg)
	if err != nil {
		log.Fatalf("Send telegram error: %v\n", err)
	}
}
