package main

import (
	"context"
	"fmt"
	"time"
)

func cookPho(ctx context.Context, ch chan<- string) {
	fmt.Println("Start cook Pho")
	select {
	case <-time.After(1 * time.Second):
		ch <- "Pho da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau pho")
		return
	}
}

func cookMi(ctx context.Context, ch chan<- string) {
	fmt.Println("Start cook Mi")
	select {
	case <-time.After(2 * time.Second):
		ch <- "Mi da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau mi")
		return
	}
}

func main() {
	chPho := make(chan string)
	chMi := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	go cookPho(ctx, chPho)
	go cookMi(ctx, chMi)

	for range 2 {
		select {
		case msg := <-chPho:
			fmt.Println("Nhan duoc: ", msg)
		case msg := <-chMi:
			fmt.Println("Nhan duoc: ", msg)
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}

	}
}
