package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type LogLevels struct {
	debug string
	info  string
	warn  string
	error string
}

type TgLogger struct {
	tgBot  *tgbotapi.BotAPI
	chatId int64 // todo: use array
	levels *LogLevels
}

func (logger *TgLogger) send(msg string) error {
	var err error
	tgMsg := tgbotapi.NewMessage(logger.chatId, msg)
	_, err = logger.tgBot.Send(tgMsg)
	return err
}

func (logger *TgLogger) debug(msg string) {
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.debug, msg))
}
func (logger *TgLogger) info(msg string) {
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.info, msg))
}
func (logger *TgLogger) warn(msg string) {
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.warn, msg))
}
func (logger *TgLogger) error(msg string) {
	_ = logger.send(fmt.Sprintf("%v\n%v", logger.levels.error, msg))
}
