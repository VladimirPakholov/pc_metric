package app

import (
	"fmt"
	"pc_metric/internal/db/repository"
	"pc_metric/internal/logger"
	"pc_metric/internal/metrics/cpu"
	net "pc_metric/internal/metrics/net_int"
	"pc_metric/internal/metrics/ram"

	"time"
)

func Start(workTime, interval time.Duration, repo *repository.Repository) {
	d := time.NewTimer(workTime)
	defer d.Stop()

	i := time.NewTicker(interval)
	defer i.Stop()

	logger.SystemMessage("=== Start getting CPU & RAM metric ===")
	logger.SystemMessage("=== Initialization... The network interface speed will be available in 10 seconds ===")

	for {
		select {
		case <-i.C:
			_, _, _, netMsg, err := net.NetMetric()
			if err != nil {
				fmt.Println("Error", err)

			}

			la := cpu.GetLoadAverage()
			r := ram.GetMemInfo()

			message := fmt.Sprintf(logger.LogMessage, la.Load1, la.Load5, la.Load15, r[0], r[1], r[2], netMsg)

			err = repo.AddMetricDB(logger.TimeStamp(), message)
			if err != nil {
				logger.SystemMessage("DB insert error: " + err.Error())
			}

			logger.LogMetric(message)

		case <-d.C:
			logger.SystemMessage("\n=== END ===")
			fmt.Println("Exit")
			return
		}
	}
}
