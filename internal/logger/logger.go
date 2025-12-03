package logger

import (
	"fmt"

	"os"
	"time"
)

var (
	fileName = "app_log.txt"
	file     *os.File
)

func InitLogger() error {

	var err error

	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	return nil
}

func LogMetric(message string) {

	if file != nil {
		//В файл добавляем временную метку
		timeTamp := time.Now().Format("2006-01-02 15:04:05")
		fileMessage := fmt.Sprintf("[%s] %s\n", timeTamp, message)

		file.WriteString(fileMessage) //запись строки в файл
	}
	//обновляемая строка в консоли
	fmt.Printf("\r%s", message)
}

// закрываем файл
func Close() {
	if file != nil {
		file.Close()
	}
}

// Статичные сообщения в консоль и файл
func SystemMessage(message string) {
	if file != nil {
		file.WriteString(message + "\n")
	}
	fmt.Println(message)
}
