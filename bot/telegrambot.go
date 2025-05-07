package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/VladShuisky/vodokanalbot/parsing"
	"github.com/VladShuisky/vodokanalbot/utils"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false // Включаем логирование

	log.Printf("Бот %s успешно запущен", bot.Self.UserName)

	// Настраиваем канал обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обрабатываем входящие сообщения
	for update := range updates {
		if update.Message == nil { // Игнорируем не-сообщения
			continue
		}
		var msg tgbotapi.MessageConfig 
		if update.Message.Text == "/get_last_info" {
			htmlFromVodokanal := parsing.GetHtmlDataFromVodokanal()
			targetTexts := parsing.ExtractText(htmlFromVodokanal)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, targetTexts[1])
			msg.ReplyToMessageID = update.Message.MessageID
		} else if update.Message.Text == "/db_healthcheck" {
			check := utils.CheckDbConnect()
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, check)
			msg.ReplyToMessageID = update.Message.MessageID
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
		}

		// Отправляем сообщение
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}