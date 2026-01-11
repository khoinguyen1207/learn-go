package main

import (
	"context"
	"sync"
	"time"

	"github.com/khoinguyen/learn-go/os-monitor/processer"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go processer.RunMonitor(ctx, &wg)

	time.Sleep(60 * time.Second)
	cancel()
	wg.Wait()
}
