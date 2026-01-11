package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)

	if err != nil {
		return "N/A"
	}

	return fmt.Sprintf("%.2f%%", percent[0])
}

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	monitor := CPUMonitor{}

	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			cpuPercent := monitor.Check(ctx)
			fmt.Println(cpuPercent)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	wg.Add(1)
	go RunMonitor(ctx, &wg)
	wg.Wait()
}
