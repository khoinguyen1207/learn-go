package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) Name() string {
	return "CPU"
}

func (m *CPUMonitor) Check(ctx context.Context) string {
	percents, err := cpu.PercentWithContext(ctx, 1*time.Second, false)

	if err != nil || len(percents) == 0 {
		return "N/A"
	}

	return fmt.Sprintf("%.2f%%", percents[0])
}
