package bot

import (
	// "fmt"
	"log"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/VladShuisky/vodokanalbot/scheduler"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false // Включаем логирование

	log.Printf("Бот %s успешно запущен", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go scheduler.StartScheduler(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := HandleUpdate(update)

		// Отправляем сообщение
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}