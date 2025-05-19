package main

import (
	"github.com/VladShuisky/vodokanalbot/bot"
	"github.com/VladShuisky/vodokanalbot/database"
)

func main() {
// Создаем бота
	// err := utils.LoadEnv(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file:", err)
	// } //закомментено, чтобы заработало на railway
	database.MakeAutomigrations()
	bot.StartBot()

}