package telegramService

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"face-id-detection/config/env"
	bardService "face-id-detection/internal/modules/bard/services"
	"face-id-detection/pkg/logger"
)

func Send() {
	bot, err := tgbotapi.NewBotAPI(env.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	logger.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var responseMsg string
	for update := range updates {
		if update.Message != nil {
			responseMsg, err = bardService.SendAsk(update.Message.Text)
			if err != nil {
				responseMsg = err.Error()
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMsg)
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}
