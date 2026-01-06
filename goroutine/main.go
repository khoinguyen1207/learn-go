package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id+1)
	time.Sleep(1 * time.Second)
}
func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := range 4 {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()

	time.Sleep(1 * time.Second)

	fmt.Println("Looped fully")

	fmt.Printf("✅ Main function completed. Elapsed time: %s\n", time.Since(start))
}
