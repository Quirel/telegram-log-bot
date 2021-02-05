package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// LevelLabels specifies labels for different log levels
type LevelLabels struct {
	debug string
	info  string
	warn  string
	error string
}

// TgLogger allows to send logs into chat with telegram bot
// todo: use array of chatId
type TgLogger struct {
	TgBot  *tgbotapi.BotAPI
	chatId int64
	levels *LevelLabels
	name   string
	level  string
}

// NewLogger creates new TgLogger
func NewLogger(level string, tgToken string, tgChatId int64) (*TgLogger, error) {
	tgBot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}

	return &TgLogger{
		TgBot:  tgBot,
		chatId: tgChatId,
		level:  level,
		levels: &LevelLabels{
			debug: "DEBUG:",
			info:  "INFO:",
			warn:  "WARN:",
			error: "ERROR:",
		},
	}, nil
}

// SetName sets name for logger displayed in message
func (logger *TgLogger) SetName(name string) {
	logger.name = name
}

// send sends message. Can return error
// todo: handle array of chatId in loop
func (logger *TgLogger) send(msg string) error {
	var err error
	if logger.name != "" {
		msg = fmt.Sprintf("%v, %v", logger.name, msg)
	}
	tgMsg := tgbotapi.NewMessage(logger.chatId, msg)
	_, err = logger.TgBot.Send(tgMsg)
	return err
}

// Log sends simple message
func (logger *TgLogger) Log(msg string) {
	_ = logger.send(msg)
}

// Debug sends debug message, depends on log level
func (logger *TgLogger) Debug(msg string) {
	if logger.level != "debug" {
		return
	}
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.debug, msg))
}

// Info sends info message, depends on log level
func (logger *TgLogger) Info(msg string) {
	if logger.level == "warn" || logger.level == "error" {
		return
	}
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.info, msg))
}

// Warn sends warn message, depends on log level
func (logger *TgLogger) Warn(msg string) {
	if logger.level == "error" {
		return
	}
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.warn, msg))
}

// Error sends error message
func (logger *TgLogger) Error(msg string) {
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.error, msg))
}
