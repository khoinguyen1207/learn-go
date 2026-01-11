package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/khoinguyen/learn-go/os-monitor/models"
	"github.com/khoinguyen/learn-go/os-monitor/monitors"
	"github.com/khoinguyen/learn-go/os-monitor/processer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	systemCh := make(chan models.SystemStats)
	monitors := []models.Monitor{
		&monitors.CPUMonitor{},
		&monitors.MEMMonitor{},
	}
	for _, monitor := range monitors {
		fmt.Println("Starting monitor:", monitor.Name())
		wg.Add(1)
		go processer.RunMonitor(ctx, &wg, systemCh, monitor)
	}

	go func() {
		for stat := range systemCh {
			fmt.Println(stat.Name, "Usage:", stat.UsagePercent)
		}
	}()

	time.Sleep(60 * time.Second)
	cancel()
	wg.Wait()
	close(systemCh)
}
