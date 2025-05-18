package scheduler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/VladShuisky/vodokanalbot/broadcast"
)

func StartScheduler(bot *tgbotapi.BotAPI) {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Minute().Do(broadcast.DoBroadcastToTelegram, bot, "Сообщение каждую минуту!")
	if err != nil {
		log.Fatal(err)
	}
	s.StartAsync()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	s.Stop()
}