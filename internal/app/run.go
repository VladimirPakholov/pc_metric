package app

import (
	"context"
	"fmt"
	"pc_metric/internal/logger"
	"pc_metric/internal/metrics/cpu"
	net "pc_metric/internal/metrics/net_int"
	"pc_metric/internal/metrics/ram"
	"pc_metric/internal/service"

	"time"
)

func Start(ctx context.Context, workTime, interval time.Duration, repo service.MetricRepository) {
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
				//fmt.Println("Error", err)
				logger.SysLogger.Error("network metric failed", "error", err)
				continue
			}

			la := cpu.GetLoadAverage()
			r := ram.GetMemInfo()

			message := fmt.Sprintf(logger.LogMessage, la.Load1, la.Load5, la.Load15, r[0], r[1], r[2], netMsg)

			//save to DB
			err = repo.AddMetric(logger.TimeStamp(), message)
			if err != nil {
				//logger.SystemMessage("DB insert error: " + err.Error())
				logger.SysLogger.Error("DB insert failed", "error", err)

			}
			//refresh live-data metric
			logger.LogMetric(message)

		case <-d.C:
			logger.SystemMessage("\n=== END ===")
			fmt.Println("Exit")
			return
		case <-ctx.Done():
			logger.SysLogger.Info("metrics stopped")
			return
		}
	}
}
