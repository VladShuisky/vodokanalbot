package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TelegramRecipient struct {
    gorm.Model        // Добавляет поля ID, CreatedAt, UpdatedAt, DeletedAt
    TelegramChatId  int64 `gorm:"unique;not null"`
    Data JSONB `gorm:"type:jsonb"`
}

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan Преобразует значение из базы данных в JSONB
func (j *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}

func GetAllRecipientsIDs() ([]int64, error) {
    db := GetDb()
    var IDs []int64

    // Пробуем получить данные
    result := db.Model(&TelegramRecipient{}).Pluck("telegram_chat_id", &IDs)

    // Если произошла ошибка (не связанная с отсутствием данных)
    if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("ошибка при получении ChatIDs: %w", result.Error)
    }

    // Если записей нет, возвращаем пустой слайс (не nil) и nil в качестве ошибки
    if len(IDs) == 0 {
        return []int64{}, nil
    }

    // Возвращаем найденные ID и nil (нет ошибки)
    return IDs, nil
}