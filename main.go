package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// init loads .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

func main() {
	tgToken := os.Getenv("TELEGRAM_LOG_BOT_TOKEN")
	tgChatId := os.Getenv("TELEGRAM_LOG_BOT_CHAT_ID")
	tgChatIdInt, _ := strconv.ParseInt(tgChatId, 10, 64)

	tgLogger, err := NewLogger("debug", tgToken, tgChatIdInt)
	if err != nil {
		log.Fatalf("error creating telegram connection: %v\n", err)
	}
	tgLogger.SetName("APP_NAME")
	tgLogger.Debug("Hello!")
}
