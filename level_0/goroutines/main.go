package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// It is a function that runs on the Go runtime.
// It helps address concurrency and async flow requirements.

// Goroutines allow you to start up and run other threads of execution
// concurrently within your program.

// Channels are used to communicate between goroutines.
// It is a typed conduit through which you can send and receive values with
// the channel operator: <-.

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMsg(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {

	sendMsg("hello") // sync

	go sendMsg("test1") // async
	go sendMsg("test2") // async
	go sendMsg("test3") // async

	sendMsg("main") // sync

	time.Sleep(2 * time.Second)

	// The sendMsg function is called synchronously and asynchronously.

	// The sendMsg function is called synchronously when the sendMsg function
	// is called without the go keyword.

	// The sendMsg function is called asynchronously when the sendMsg function
	//  is called with the go keyword.

	// How Does a Goroutine Work?

	// When the sendMsg function is called with the go keyword, the main function
	// will not wait for the sendMsg function to finish executing before it
	// continues to the next line of code and will return immediately after the
	// sendMsg function is called.

	// Otherwise, the function is called synchronously, and the main function
	// will wait for the sendMsg function to finish executing before it continues
	// to the next line of code.

	// The order of the output when you run the above example will differ from
	// the order of the code because the three goroutine all run concurrently
	// and since the functions pause for a period of time, the order which they
	// wake will differ and be outputted.

}
