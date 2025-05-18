package main

import (
	"log"

	// "fmt"
	// "log"
	// "os"

	"github.com/VladShuisky/vodokanalbot/bot"
	"github.com/VladShuisky/vodokanalbot/database"
	// "github.com/VladShuisky/vodokanalbot/parsing"
	"github.com/VladShuisky/vodokanalbot/utils"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	// "github.com/VladShuisky/vodokanalbot/utils"
	// "github.com/VladShuisky/vodokanalbot/bot"
	// "log"
	// "github.com/jackc/pgx/v5"
)

func main() {
// Создаем бота
	// err := utils.LoadEnv(".env")
	// if err != nil {
		// log.Fatal("Error loading .env file:", err)
	// } //закомментено, чтобы заработало на railway
	database.MakeAutomigrations()
	bot.StartBot()

}