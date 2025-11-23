package logger

import (
	"fmt"

	"os"
	"time"
)

var (
	FileName = "app_log.txt"
	File     *os.File
)

func InitLogger() error {

	var err error

	File, err = os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	return nil
}

func LogMetric(message string) {

	if File != nil {
		//В файл добавляем временную метку
		timeTamp := time.Now().Format("2006-01-02 15:04:05")
		fileMessage := fmt.Sprintf("[%s] %s\n", timeTamp, message)

		File.WriteString(fileMessage)
	}
	//обновляемая строка в консоли
	fmt.Printf("\r%s", message)
}

func Close() {
	if File != nil {
		File.Close()
	}
}
