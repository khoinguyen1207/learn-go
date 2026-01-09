package main

import (
	"fmt"
	"sync"
	"time"
)

func task(i int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	fmt.Println("Task started")
	time.Sleep(1 * time.Second)
	ch <- "42"
}

func main() {
	var wg sync.WaitGroup
	var ch = make(chan string)
	// go func() {
	// 	defer close(ch)
	// 	ch <- 1
	// 	ch <- 2
	// 	ch <- 3
	// }()

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-ch)
	// }

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// =====================
	start := time.Now()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go task(i, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Printf("✅ Main function completed. Elapsed time: %s\n", time.Since(start))
}
