package main

import (
	"fmt"
	"time"
)

var weekday string

// In Go, the predefined init() function sets off a piece of code to run before
// any other part of your package. This code will execute as soon as the
// package is imported, and can be used when you need your application to
// initialize in a specific state, such as when you have a specific configuration
// or set of resources with which your application needs to start. It is also
// used when importing a side effect, a technique used to set the state of a
// program by importing a specific package. This is often used to register one
// package with another to make sure that the program is considering the correct
// code for the task.
func init() {
	weekday = time.Now().Weekday().String()
}

func main() {
	// one very diffrent thing is happening here
	// and keep in mind that in a module there can be diffrent packages
	// with a main.go file as it is possible in go to have diffrent
	// executable program in a single module
	// also you can have diffrent files inside functionInGo folder
	// the only constarint is,they should have same package name
	// as in go,all files inside a folder should have same package name

	// even though this function is defined in other file but because of the
	// same package name,it can be used without importing
	// the only thing is instead of go run main.go you need to do go run .
	// as you need to compile all the file and not just main.go

	variadic("navneet", "ankit", "dev", "ritesh bhaiya", "rohit", "rohit mote", "jai", "isham")

	fmt.Printf("today is %s ", weekday)
}
