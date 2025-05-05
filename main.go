package main

import (
	// "log"
	// "github.com/VladShuisky/vodokanalbot/utils"
	"github.com/VladShuisky/vodokanalbot/bot"
)


func main() {
	// Создаем бота
	// err := utils.LoadEnv(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file:", err)
	// } //закомментено, чтобы заработало на railway
	bot.StartBot()
}