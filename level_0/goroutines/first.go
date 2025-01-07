package main

import (
	"fmt"
	"time"
)

// this is after main means it is second chapter

//  goroutines are lightweight threads managed by the Go runtime.
// They enable concurrent execution of functions,
// Goroutines are much smaller and more efficient than operating system threads.

// They consume less memory (around 2KB initial stack size) and are multiplexed
//  onto a small number of OS threads by the Go runtime.

// Goroutines allow multiple functions to run simultaneously without blocking
// each other.
// This is not the same as parallelism (using multiple CPU cores), but it can
// be used to achieve it.
// Goroutines execute independently, and their scheduling is handled by the Go
//  runtime’s scheduler.

// The Go scheduler can interrupt goroutines and switch between them,
// ensuring fairness and efficiency.

// A goroutine is created by prepending the go keyword to a function call.

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

// if you call it normally,it is just a syncronous call
// but when you call it like this $ go printMessage(),then a goroutine
// will be created

// Behaviour

// When a goroutine is launched, the calling function does not wait for it to finish.
// The main function runs in its own goroutine (called the "main goroutine").
// If the main goroutine exits, all other goroutines are terminated immediately.
// so we need to wait for completion of all the goroutine

// Ways to wait

// 1. time.sleep()=> not ideal
// 2. Synchronization Techniques
//            => WaitGroup (from sync package).
//            => Channels (Go's built-in concurrency mechanism).
