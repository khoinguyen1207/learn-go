package processer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/khoinguyen/learn-go/os-monitor/monitors"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	cpuMonitor := monitors.CPUMonitor{}
	memMonitor := monitors.MEMMonitor{}

	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			cpuPercent := cpuMonitor.Check(ctx)
			memPercent := memMonitor.Check(ctx)
			fmt.Println("CPU:", cpuPercent, "MEM:", memPercent)
		}
	}
}
