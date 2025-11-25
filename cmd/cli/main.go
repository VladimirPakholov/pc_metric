package main

import (
	"fmt"

	"time"

	"pc_metric/internal/logger"
	"pc_metric/metrics/cpu"
	net "pc_metric/metrics/net_int"
	"pc_metric/metrics/ram"
)

func main() {
	startApp := time.Now()
	durationWork := 1 * time.Minute

	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	//net.DebugInterfaces()

	logger.SystemMessage("=== Start getting CPU & RAM metric ===")
	logger.SystemMessage("=== Initialization... The network interface speed will be available in 10 seconds ===")

	for time.Since(startApp) < durationWork {
		_, _, _, netMsg, err := net.NetMetric()
		if err != nil {
			fmt.Println("Error", err)

		}

		la := cpu.GetLoadAverage()
		r := ram.GetMemInfo()

		message := fmt.Sprintf("Load average is: 1 min: %.2f, 5 min: %.2f, 15 min: %.2f | RAM: %v/%vGB (%vGB free) | NET: %s ", la.Load1, la.Load5, la.Load15, r[0], r[1], r[2], netMsg)

		logger.LogMetric(message)

		time.Sleep(1 * time.Second)
	}

	logger.SystemMessage("\n=== END ===")

}
