package main

import (
	"flag"
	"fmt"
	"os"
	"pc_metric/internal/database"
	"pc_metric/internal/logger"
	"pc_metric/internal/metrics/cpu"
	net "pc_metric/internal/metrics/net_int"
	"pc_metric/internal/metrics/ram"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	//Загрузка файла переменных окружения
	if err := godotenv.Load(); err != nil {
		fmt.Println("File .env not found")
		os.Exit(1)
	}
	cfg := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
		SSLmode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()
	err = database.NewTable(db)
	if err != nil {
		os.Exit(1)
	}

	//читаем строку из файла с переменными
	defaultTimeWorkStr := os.Getenv("DEFAULT_TIME_WORK")

	defaultTimeWorkT, err := time.ParseDuration(defaultTimeWorkStr)
	if err != nil {
		fmt.Printf("Parse error time from .env: %v. Set default time 1m\n", err)
		defaultTimeWorkT = 1 * time.Minute
	}
	defaultTimeMetricStr := os.Getenv("DEFAULT_TIME_GET_METRIC")

	defaultTimeMetricT, err := time.ParseDuration(defaultTimeMetricStr)
	if err != nil {
		fmt.Printf("Error parse getting time metric: %v. Set default time 1s\n", err)
		defaultTimeMetricT = 1 * time.Second
	}

	//стандартное  время работы приложения 1 минута, если не передано иное
	timeWorkDefault := flag.Duration("d", defaultTimeWorkT, "Default time for working app - 1 min") // -d 2m, -d 30s, -d 1h
	//стандартное время сбора метрик 1 секунда, если не передано иное
	timeMetricDefault := flag.Duration("i", defaultTimeMetricT, "Default time for metrics monitoring - 1 second")
	flag.Parse()

	if *timeWorkDefault <= 0 {
		fmt.Println("Duration work time must be > 0")
		os.Exit(1)
	}
	if *timeMetricDefault <= 0 {
		fmt.Println("Duration get metric must be > 0")
		os.Exit(1)
	}

	ticker := time.NewTicker(*timeMetricDefault) // как часто собирать метрики
	defer ticker.Stop()

	timer := time.NewTimer(*timeWorkDefault) // общее время работы программы
	defer timer.Stop()

	logger.SystemMessage("=== Start getting CPU & RAM metric ===")
	logger.SystemMessage("=== Initialization... The network interface speed will be available in 10 seconds ===")

	for {
		select {
		case <-ticker.C:

			_, _, _, netMsg, err := net.NetMetric()
			if err != nil {
				fmt.Println("Error", err)

			}

			la := cpu.GetLoadAverage()
			r := ram.GetMemInfo()

			message := fmt.Sprintf(logger.LogMessage, la.Load1, la.Load5, la.Load15, r[0], r[1], r[2], netMsg)
			err = database.InsertLogMetric(db, logger.TimeStamp(), message)
			if err != nil {
				logger.SystemMessage("DB insert error: " + err.Error())
			}

			logger.LogMetric(message)

		case <-timer.C:
			logger.SystemMessage("\n=== END ===")
			fmt.Println("Exit")
			return
		}

	}

}
