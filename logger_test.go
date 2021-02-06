package tglogger_test

import (
	tglog "github.com/quirel/telegram-logger"
	"os"
	"strconv"
	"testing"
)

var TestToken = os.Getenv("TEST_TOKEN")
var TestChatId, _ = strconv.ParseInt(
	os.Getenv("TEST_CHAT_ID"), 10, 64)

func TestAssertEnv(t *testing.T) {
	if TestToken == "" {
		t.Error("Specify TEST_TOKEN in environment")
		t.Fail()
	}
	if TestChatId == 0 {
		t.Error("Specify TEST_CHAT_ID in environment")
		t.Fail()
	}
}

func TestNewLogger(t *testing.T) {
	_, err := tglog.NewLogger("Debug", TestToken, []int64{TestChatId})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestLog(t *testing.T) {
	logger, _ := tglog.NewLogger("Debug", TestToken, []int64{TestChatId})
	err := logger.Send("TestSend", TestChatId)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSetLabels(t *testing.T) {
	labels := tglog.LevelLabels{
		Debug: "dbg", Info: "inf", Warn: "wrn", Error: "err",
	}
	logger, _ := tglog.NewLogger("Debug", TestToken, []int64{TestChatId})

	if logger.Labels.Debug != "DEBUG:" {
		t.Error("Default Labels.Debug should be 'DEBUG:'")
		t.Fail()
	}
	logger.SetLabels(&labels)
	if logger.Labels.Debug != "dbg" {
		t.Error("Labels.Debug should be set to 'dbg'")
		t.Fail()
	}
}
