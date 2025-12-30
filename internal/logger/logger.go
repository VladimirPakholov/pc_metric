package logger

import (
	"fmt"
	"time"
)

var (
	LogMessage = "Load average is: 1 min: %.2f, 5 min: %.2f, 15 min: %.2f | RAM: %v/%vGB (%vGB free) | NET: %s "
)

func LogMetric(message string) {
	//обновляемая строка в консоли
	fmt.Printf("\r%s", message)
}

// Статичные сообщения в консоль и файл
func SystemMessage(message string) {
	fmt.Println(message)
}

// Временная метка для БД
func TimeStamp() time.Time {
	return time.Now()
}
