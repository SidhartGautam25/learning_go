package main

import "fmt"

// Channels in Go are a mechanism for goroutines to communicate and synchronize
//  by sending and receiving values.

// A channel is a typed conduit through which you can send and receive values
//  between goroutines.

// Think of a channel as a pipe: one goroutine sends data into the pipe,
// and another goroutine receives it.

// Channel Operations
// Send: channel <- value
// Receive: value := <-channel
// Close: close(channel) (to indicate no more values will be sent)

// Sending to or receiving from an unbuffered channel blocks the
// sending/receiving goroutine until the other side is ready.
// This makes channels useful for synchronization.

// declaring a channel
// var c chan int // Declare a channel for int type
// c = make(chan int) // Initialize the channel

// Or simply:
// c := make(chan int) // Declare and initialize

func Example() {
	c := make(chan int) // Create a channel

	go func() {
		c <- 42 // Send a value into the channel
	}()

	value := <-c // Receive a value from the channel
	fmt.Println("Received:", value)
	//output
	// Received: 42
}

func PrintNumbers(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i // Send data to the channel
	}
	close(c) // Close the channel after sending all data
}

func LearningChannels() {
	c := make(chan int)

	go PrintNumbers(c) // Start a goroutine

	// Receive data from the channel
	for num := range c {
		fmt.Println(num)
	}
}

// Unbuffered Channels:

// Data transfer happens only when both sender and receiver are ready.
// Use case: Synchronization between goroutines.
func Unbuffered() {
	c := make(chan int)

	go func() {
		c <- 10 // Send will block until main goroutine receives
	}()

	fmt.Println(<-c) // Receive from channel
}

// Buffered Channels:

// Can hold a limited number of values without a receiver.
// Sending blocks only when the buffer is full.
// Receiving blocks only when the buffer is empty.

func Buffered() {

	c := make(chan int, 2) // Create a buffered channel with capacity 2

	c <- 1
	c <- 2 // Buffered, so no need for immediate receiver

	fmt.Println(<-c) // Receive the first value
	fmt.Println(<-c) // Receive the second value
}

// Closing a Channel
// Closing a channel signals that no more values will be sent.
// Receivers can still receive any remaining values.
// Sending to a closed channel causes a panic.

func Closing() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	close(c) // Close the channel

	for value := range c {
		fmt.Println(value)
	}
}

// using channels for syncronization
func Worker(done chan bool) {
	fmt.Println("Working...")
	done <- true // Notify main goroutine
}

func UsingWorker() {
	done := make(chan bool)
	go Worker(done)
	<-done // Wait for notification
	fmt.Println("Done")
}

// Using Channels for Communication
func Producer(c chan int) {
	for i := 1; i <= 5; i++ {
		c <- i // Send numbers to the channel
	}
	close(c) // Close the channel after sending
}

func UsingProducer() {

	c := make(chan int)

	go Producer(c)

	for value := range c { // Receive until the channel is closed
		fmt.Println(value)
	}
}

// select keyword
