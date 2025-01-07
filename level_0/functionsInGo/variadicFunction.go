package main

import "fmt"

// A variadic function is a function that accepts zero, one, or more values
//  as a single argument.

// this is a variadic function
func variadic(names ...string) {

	// The _ is the blank identifier in Go.
	// It is used when you want to ignore a value.
	// Specifically in this case, _ is used to ignore the index returned by the range keyword.
	for _, name := range names {
		fmt.Printf("Hello %s\n ", name)
	}
}
