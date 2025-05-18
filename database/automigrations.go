package database

func MakeAutomigrations() {
	db := GetDb()
	db.AutoMigrate(&TelegramRecipient{})
}