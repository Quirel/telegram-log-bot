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

	tgChatId := os.Getenv("TELEGRAM_LOG_BOT_CHAT_ID")
	tgChatIdInt, _ := strconv.ParseInt(tgChatId, 10, 64)

	// todo: handle channel
	tgLogger := TgLogger{
		tgBot:  tgBot,
		chatId: tgChatIdInt,
		levels: &LogLevels{
			debug: "DEBUG:",
			info:  "INFO:",
			warn:  "WARN:",
			error: "ERROR:",
		},
	}

	tgLogger.debug("Hello!")
	tgLogger.info("Hello!")
	tgLogger.warn("Hello!")
	tgLogger.error("Hello!")
}
