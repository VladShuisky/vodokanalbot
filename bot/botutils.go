package bot

import (
	"fmt"
	"time"
)

func GetStartMessage() string {
	now := time.Now()
	dateStr := now.Format("02.01.2006")
	message := fmt.Sprintf("Приветствую! Вот список команд: \n/start\n/get_last_info\n/date %s\n/db_healthcheck", dateStr)
	return message
}