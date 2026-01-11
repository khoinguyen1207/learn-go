package processer

import (
	"context"
	"sync"
	"time"

	"github.com/khoinguyen/learn-go/os-monitor/models"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, systemCh chan<- models.SystemStats, monitor models.Monitor) {
	defer wg.Done()
	m := monitor

	timer := time.NewTicker(2 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			usagePercent := m.Check(ctx)
			systemCh <- models.SystemStats{
				Name:         m.Name(),
				UsagePercent: usagePercent,
			}
		}
	}
}
