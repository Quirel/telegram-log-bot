// Package provides primitives for sending logs
// into telegram chat
package tglogger

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

// LevelLabels specifies labels for different log levels
type LevelLabels struct {
	debug string
	info  string
	warn  string
	error string
}

// TgLogger allows to send logs into chat with telegram bot
// level - debug | info | warn | error
type TgLogger struct {
	TgBot      *tgbotapi.BotAPI
	chatIdList []int64
	levels     *LevelLabels
	name       string
	level      string
}

// NewLogger creates new TgLogger
func NewLogger(level string, tgToken string, tgChatIds []int64) (*TgLogger, error) {
	tgBot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}

	return &TgLogger{
		TgBot:      tgBot,
		chatIdList: tgChatIds,
		level:      level,
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

// Send sends message. Can return error
func (logger *TgLogger) Send(msg string, chatId int64) error {
	var err error
	if logger.name != "" {
		msg = fmt.Sprintf("%v, %v", logger.name, msg)
	}
	tgMsg := tgbotapi.NewMessage(chatId, msg)
	_, err = logger.TgBot.Send(tgMsg)
	return err
}

// sendWithoutError sends message, suppressing error
func (logger *TgLogger) sendWithoutError(msg string, chatId int64) {
	_ = logger.Send(msg, chatId)
}

// SendMultiple sends message to multiple chats
func (logger *TgLogger) SendMultiple(msg string) {
	var wg sync.WaitGroup
	for _, id := range logger.chatIdList {
		wg.Add(1)
		go func(msg string, id int64) {
			defer wg.Done()
			logger.sendWithoutError(msg, id)
		}(msg, id)
	}
	wg.Wait()
}

// Log sends simple message
func (logger *TgLogger) Log(msg string) {
	logger.SendMultiple(msg)
}

// Debug sends debug message, depends on log level
func (logger *TgLogger) Debug(msg string) {
	if logger.level != "debug" {
		return
	}
	logger.SendMultiple(fmt.Sprintf("%v\n%v", logger.levels.debug, msg))
}

// Info sends info message, depends on log level
func (logger *TgLogger) Info(msg string) {
	if logger.level == "warn" || logger.level == "error" {
		return
	}
	logger.SendMultiple(fmt.Sprintf("%v\n%v", logger.levels.info, msg))
}

// Warn sends warn message, depends on log level
func (logger *TgLogger) Warn(msg string) {
	if logger.level == "error" {
		return
	}
	logger.SendMultiple(fmt.Sprintf("%v\n%v", logger.levels.warn, msg))
}

// Error sends error message
func (logger *TgLogger) Error(msg string) {
	logger.SendMultiple(fmt.Sprintf("%v\n%v", logger.levels.error, msg))
}
