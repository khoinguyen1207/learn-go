package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MEMMonitor struct {
}

func (m *MEMMonitor) Name() string {
	return "MEM"
}

func (m *MEMMonitor) Check(ctx context.Context) string {
	virtualMem, err := mem.VirtualMemoryWithContext(ctx)

	if err != nil {
		return "N/A"
	}

	return fmt.Sprintf("%.2f%%", virtualMem.UsedPercent)
}
