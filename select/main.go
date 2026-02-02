package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	strChan := make(chan string)
	numChan := make(chan int)

	go func() {
		for {
			strChan <- "Hello"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for {
			numChan <- 1
			time.Sleep(2 * time.Second)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Receive from whichever channel is ready first
	for {
		select {
		case msg := <-strChan:
			fmt.Println("string:", msg)

		case num := <-numChan:
			fmt.Println("number:", num)

		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
			return
		}
	}
}
