package utils

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	// Получаем общее количество секунд
	totalSeconds := int64(d.Seconds())
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	// Форматируем строку
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
