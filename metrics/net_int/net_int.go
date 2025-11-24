package net_int

import (
	"github.com/shirou/gopsutil/v4/net"
)

var (
	IntName  string
	SendByte uint64
	RecByte  uint64
)

func NetMetric() {
	metric, _ := net.IOCounters(true)

	for _, stat := range metric {
		if stat.Name == "en0" {
			IntName = stat.Name
		} else if stat.BytesSent > 0 && stat.BytesRecv > 0 {
			SendByte = stat.BytesSent
			RecByte = stat.BytesRecv
		}

		switch stat.Name {

		case "en0":

		}

	}
}
