package bot

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Включаем логирование

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

		// Создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		// Отправляем сообщение
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}