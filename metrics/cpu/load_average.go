package cpu

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/load"
)

func GetLoadAverage() *load.AvgStat {

	avg, err := load.Avg()
	if err != nil {
		fmt.Println("Error to get LoadAverage CPU")
	}

	return avg
}
