package main

import (
	"fmt"
	"sync"
)

// Using sync.WaitGroup

func goroutineFunction(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Println(msg)
}

func second() {
	var wg sync.WaitGroup

	wg.Add(2) // Add two goroutines to wait for

	go goroutineFunction("Goroutine 1", &wg)
	go goroutineFunction("Goroutine 2", &wg)

	wg.Wait() // Block until all goroutines finish
	fmt.Println("All goroutines completed")
}

// if you run second(),output will be
// Goroutine 1
// Goroutine 2
// All goroutines completed
