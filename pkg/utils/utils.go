package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseEnvInt(envKey string) (int, error) {
	value, err := strconv.Atoi(os.Getenv(envKey))
	if err != nil {
		return 0, fmt.Errorf("invalid %s :%w", envKey, err)
	}
	return value, nil
}

func ValidTime(envKey string) (time.Duration, error) {
	time, err := time.ParseDuration(os.Getenv(envKey))
	if err != nil {
		return 0, fmt.Errorf("invalid %s :%w", envKey, err)
	}
	return time, nil
}

// LoadEnv вручную читает .env и загружает переменные в окружение. Аналог сторонней: godotenv
func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err // .env не найден или не читается
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Пропускаем комментарии и пустые строки
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Разделяем по первому знаку "="
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // невалидная строка
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		// Удалим кавычки если есть
		val = strings.Trim(val, `"'`)

		// Устанавливаем переменную окружения
		_ = os.Setenv(key, val)
	}

	return scanner.Err()
}
