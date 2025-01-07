package main

import "fmt"

func main() {
	// In Go, defer is a powerful feature that schedules a function call to be
	// executed after the surrounding function returns, regardless of how the
	// function exits (whether it returns normally or due to a panic).

	// defer statements are executed in Last In, First Out (LIFO) order when
	// the surrounding function returns.

	// Deferred functions are executed after the function's logic is complete but
	// before the function returns to its caller.

	// The arguments of the deferred function are evaluated at the time the
	// defer statement is executed, not when the deferred function is invoked.
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Function body")

	// output will be
	// Function body
	// Third
	// Second
	// First

	// we will learn the use cases of defer after learning basic stuffs
}
