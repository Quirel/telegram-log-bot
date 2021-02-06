# Telegram logger

Logger to send message in telegram chats.

## Example

```go
package main

import tglog "github.com/quirel/telegram-logger"

func main() {
	token := "YOUR_TELEGRAM_BOT_TOKEN"
	chatIds := []int64{123123, 456456}

	// creates logger with level 'debug'
	logger, _ := tglog.NewLogger("debug", token, chatIds)
	// set name displayed in logs
	logger.SetName("MyApp")
	
	logger.Debug("Debug message") // "MyApp, DEBUG:\nDebug message"
	logger.Info("Info message") // "MyApp, INFO:\nInfo message"
	logger.Warn("Warning message") // "MyApp, WARN:\nWarning message"
	logger.Error("Error message") // "MyApp, ERROR:\nError message"

	labels := tglog.LevelLabels{
		Debug: "dbg", Info: "inf", Warn: "wrn", Error: "err",
	}
	// Set not default labels for different log levels
	logger.SetLabels(&labels)
	logger.Debug("Debug message") // "MyApp, dbg\nDebug message"
}
```
