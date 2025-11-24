package main

import (
	"fmt"
	"time"

	"pc_metric/metrics/cpu"
	"pc_metric/metrics/ram"
	"pc_metric/pkg/logger"
)

func main() {

	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	//fmt.Println("=== Testing gopsutil ===")
	logger.SystemMessage("=== Start getting CPU & RAM metric ===")

	for i := 0; i < 25; i++ {

		la := cpu.GetLoadAverage()
		r := ram.GetMemInfo()
		message := fmt.Sprintf("Load average is: 1 min: %.2f, 5 min: %.2f, 15 min: %.2f | RAM: %v/%vGB (%vGB free)", la.Load1, la.Load5, la.Load15, r[0], r[1], r[2])

		logger.LogMetric(message)
		time.Sleep(1 * time.Second)
	}

	logger.SystemMessage("\n=== END ===")

}
