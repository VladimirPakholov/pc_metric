package main

import (
	"fmt"
	"time"

	"pc_metric/metrics/cpu"
	"pc_metric/metrics/ram"
	"pc_metric/pkg/logger"
	//"github.com/shirou/gopsutil/v4/mem"
	//"github.com/shirou/gopsutil/v4/load"
	//"github.com/shirou/gopsutil/v4/cpu"
	// "github.com/shirou/gopsutil/v4/mem"
)

func main() {
	// Тест загрузки CPU

	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	fmt.Println("=== Testing gopsutil ===")

	for i := 0; i < 25; i++ {

		la := cpu.GetLoadAverage()
		r := ram.GetMemInfo()
		message := fmt.Sprintf("Load average is: 1 min: %.2f, 5 min: %.2f, 15 min: %.2f | RAM: %v/%vGB (%vGB free)", la.Load1, la.Load5, la.Load15, r[0], r[1], r[2])

		logger.LogMetric(message)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n=== END ===")

	if logger.File != nil {
		logger.File.WriteString("\n=== END ===")
	}
}

//     // 1. Средняя загрузка
//     avg, err := load.Avg()
//     if err != nil {
//         log.Printf("Load average error: %v", err)
//     } else {
//         fmt.Printf("Load: %.2f, %.2f, %.2f\n", avg.Load1, avg.Load5, avg.Load15)
//     }

//     // 2. Информация о CPU
//     cpuInfo, err := cpu.Info()
//     if err != nil {
//         log.Printf("CPU info error: %v", err)
//     } else if len(cpuInfo) > 0 {
//         fmt.Printf("CPU: %s\n", cpuInfo[0].ModelName)
//     }

//     // 3. Память
//     memory, err := mem.VirtualMemory()
//     if err != nil {
//         log.Printf("Memory error: %v", err)
//     } else {
//         fmt.Printf("Memory: %.2f%% used\n", memory.UsedPercent)
//     }

// fmt.Println("=== All dependencies working! ===")
