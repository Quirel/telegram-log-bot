# Telegram logger

Logger to send message in telegram chats.

## Example

```go
package main

import tglog "github.com/quirel/telegram-logger"

func main() {
	token := "YOUR_TELEGRAM_BOT_TOKEN"
	var chatId int64 = 123123

	// creates logger with level 'debug'
	logger, _ := tglog.NewLogger("debug", token, []int64{chatId})
	// set name displayed in logs
	logger.SetName("MyApp")
	
	logger.Debug("Debug message")
	logger.Info("Debug message")
	logger.Warn("Debug message")
	logger.Error("Debug message")
}
```
