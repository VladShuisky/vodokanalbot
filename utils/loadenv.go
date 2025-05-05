package utils

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue // Пропускаем комментарии и пустые строки
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Некорректная строка
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value) // Устанавливаем переменную окружения
	}

	return scanner.Err()
}