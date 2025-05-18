package broadcast

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/VladShuisky/vodokanalbot/database"
)

func DoBroadcastToTelegram(bot *tgbotapi.BotAPI, messageText string) {
	chatIDs, err := database.GetAllRecipientsIDs()
	if err != nil {
		log.Fatal("Ошибка в фунции GetAllRecipientsIDS")
	}
	SendBroadcast(bot, messageText, chatIDs)
}

func SendBroadcast(bot *tgbotapi.BotAPI, messageText string, recipients []int64) {
	chatIDs := recipients
	for _, chatID := range chatIDs {
		msg := tgbotapi.NewMessage(chatID, messageText)
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Ошибка отправки в чат: %d: %v", chatID, err)
		}
	}
}