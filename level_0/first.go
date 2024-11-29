package main

import (
	"fmt"
)

func main() {
	var name string

	// Ask for the user's name
	fmt.Print("Enter your name: ");
	// Read the name input from the user
	fmt.Scanln(&name);

	// Print a greeting message
	fmt.Printf("Hello, %s! Welcome to the world of Go.\n", name);
}
