package net_int

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/net"
)

var (
	IntName string

	prevSend uint64 //для вычисления разницы
	prevRecv uint64

	initialized bool
	startTime   time.Time
	lastTime    time.Time
)

// DetectActiveInterface ищет активный сетевой интерфейс (игнорирует lo)
func AutoDetectInterface() (string, error) {
	metrics, err := net.IOCounters(true)
	if err != nil {
		return "", err
	}
	var activeInterface string
	var maxTraffic uint64

	for _, stat := range metrics {
		if stat.Name == "lo" ||
			strings.HasPrefix(stat.Name, "awdl") ||
			strings.HasPrefix(stat.Name, "llw") ||
			strings.HasPrefix(stat.Name, "bridge") ||
			strings.HasPrefix(stat.Name, "p2p") {
			continue
		}

		traffic := stat.BytesRecv + stat.BytesSent
		if traffic > maxTraffic {
			maxTraffic = traffic
			activeInterface = stat.Name
		}
	}
	if activeInterface == "" {
		return "", errors.New("Not found active interface")
	}
	return activeInterface, nil
}

// NetMetric возвращает скорость интерфейса в KB/s.
func NetMetric() (upKBps, downKBps uint64, readyNow bool, msg string, err error) {

	if IntName == "" {
		iface, err := AutoDetectInterface()
		if err != nil {
			return 0, 0, false, "", err
		}
		IntName = iface
	}

	metric, err := net.IOCounters(true)
	if err != nil {
		return 0, 0, false, "", err
	}

	for _, stat := range metric {
		if stat.Name == IntName {

			if !initialized {
				prevSend = stat.BytesSent
				prevRecv = stat.BytesRecv
				startTime = time.Now()
				lastTime = startTime
				initialized = true
				return 0, 0, false, "", nil // сразу возвращаем 0 скорости
			}

			// если прошло меньше 10 секунд — тоже возвращаем 0
			if time.Since(startTime) < 10*time.Second {
				return 0, 0, false, "Initialization...", nil
			}

			//Вычисление скорости канала

			elapsed := time.Since(lastTime).Seconds()
			if elapsed <= 0 {
				elapsed = 1
			}

			upKBps = uint64(float64(stat.BytesSent-prevSend) / 1024 / elapsed)
			downKBps = uint64(float64(stat.BytesRecv-prevRecv) / 1024 / elapsed)

			prevSend = stat.BytesSent
			prevRecv = stat.BytesRecv
			lastTime = time.Now()

			msg := fmt.Sprintf("UP: %d KB/s DOWN: %d KB/s", upKBps, downKBps)

			return upKBps, downKBps, true, msg, nil
		}

	}
	return 0, 0, false, "", errors.New("Interface not found")
}
