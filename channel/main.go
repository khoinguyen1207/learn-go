package main

import "fmt"

func main() {
	var ch = make(chan int)

	go func() {
		defer close(ch)
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-ch)
	// }

	for v := range ch {
		fmt.Println(v)
	}
}
