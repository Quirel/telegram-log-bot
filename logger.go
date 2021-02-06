// Package to send message in telegram chats.
// Package created for some log notification for pet/small projects.
// It's not suitable for something like access.log on highloaded projects.
package tglogger

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
)

// LevelLabels specifies labels for different log Labels
type LevelLabels struct {
	Debug string
	Info  string
	Warn  string
	Error string
}

// TgLogger allows to send logs into chat with telegram bot
// Level - Debug | Info | Warn | Error
type TgLogger struct {
	TgBot      *tgbotapi.BotAPI
	ChatIdList []int64
	Labels     *LevelLabels
	Name       string
	Level      string
}

// NewLogger creates new TgLogger
func NewLogger(level string, tgToken string, tgChatIds []int64) (*TgLogger, error) {
	tgBot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		return nil, err
	}

	return &TgLogger{
		TgBot:      tgBot,
		ChatIdList: tgChatIds,
		Level:      level,
		Labels: &LevelLabels{
			Debug: "DEBUG:",
			Info:  "INFO:",
			Warn:  "WARN:",
			Error: "ERROR:",
		},
	}, nil
}

// SetName sets Name for logger displayed in message
func (logger *TgLogger) SetName(name string) {
	logger.Name = name
}

func (logger *TgLogger) SetLabels(labels *LevelLabels) {
	logger.Labels = labels
}

// Send sends message. Can return Error
func (logger *TgLogger) Send(msg string, chatId int64) error {
	var err error
	if logger.Name != "" {
		msg = fmt.Sprintf("%v, %v", logger.Name, msg)
	}
	tgMsg := tgbotapi.NewMessage(chatId, msg)
	_, err = logger.TgBot.Send(tgMsg)
	return err
}

// sendWithoutError sends message, suppressing Error
func (logger *TgLogger) sendWithoutError(msg string, chatId int64) {
	_ = logger.Send(msg, chatId)
}

// Log sends simple message
func (logger *TgLogger) Log(msg string) {
	var wg sync.WaitGroup
	for _, id := range logger.ChatIdList {
		wg.Add(1)
		go func(msg string, id int64) {
			defer wg.Done()
			logger.sendWithoutError(msg, id)
		}(msg, id)
	}
	wg.Wait()
}

// Debug sends Debug message, depends on log Level
func (logger *TgLogger) Debug(msg string) {
	if logger.Level != "Debug" {
		return
	}
	logger.Log(fmt.Sprintf("%v\n%v", logger.Labels.Debug, msg))
}

// Info sends Info message, depends on log Level
func (logger *TgLogger) Info(msg string) {
	if logger.Level == "Warn" || logger.Level == "Error" {
		return
	}
	logger.Log(fmt.Sprintf("%v\n%v", logger.Labels.Info, msg))
}

// Warn sends Warn message, depends on log Level
func (logger *TgLogger) Warn(msg string) {
	if logger.Level == "Error" {
		return
	}
	logger.Log(fmt.Sprintf("%v\n%v", logger.Labels.Warn, msg))
}

// Error sends Error message
func (logger *TgLogger) Error(msg string) {
	logger.Log(fmt.Sprintf("%v\n%v", logger.Labels.Error, msg))
}
