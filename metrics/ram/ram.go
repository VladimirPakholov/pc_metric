package ram

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func GetMemInfo() []int {
	GB := 1024 * 1024 * 1024

	ram, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error get RAM value")
	}
	memStat := []int{}
	ramGbTotal := int(ram.Total) / GB
	memStat = append(memStat, ramGbTotal)
	ramGbUsed := int(ram.Used) / GB
	memStat = append(memStat, ramGbUsed)
	ramGbAvailable := int(ram.Available) / GB
	memStat = append(memStat, ramGbAvailable)
	return memStat
}
