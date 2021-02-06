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
	
	logger.Debug("Debug message")
	logger.Info("Debug message")
	logger.Warn("Debug message")
	logger.Error("Debug message")
}
```
