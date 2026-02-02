package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}


func waiter(id int, orders <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Waiter %d: Starting work...\n", id)
	
	for order := range orders {		
	
		fmt.Printf("Waiter %d: Preparing order for table %d...\n", id, order.TableNumber)
		time.Sleep(order.PrepTime)
		
		fmt.Printf("Waiter %d: Order ready for table %d!\n\n", id, order.TableNumber)
	}
}

func main() {
	start := time.Now()
	orders := []Order{
		{TableNumber: 1, PrepTime: 8 * time.Second},
		{TableNumber: 2, PrepTime: 5 * time.Second},
		{TableNumber: 3, PrepTime: 5 * time.Second},
		{TableNumber: 4, PrepTime: 3 * time.Second},
	}
	
	const numWaiters = 3
	
	jobs := make(chan Order, len(orders))
	

	var wg sync.WaitGroup
	for i := 1; i <= numWaiters; i++ {
		wg.Add(1)
		go waiter(i, jobs, &wg)
	}

	fmt.Println("Adding orders to queue...")
	for _, order := range orders {
		jobs <- order
	}
	close(jobs)
	
	wg.Wait()
	
	fmt.Println("all orders processed!")
	fmt.Printf("Time taken: %v\n", time.Since(start))
}