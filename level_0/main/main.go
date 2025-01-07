package main

//  This means that when you create a repository that you intend to share, the
//  package name should be the repository name, and the package’s source should be in
//  the root of the repository’s directory structure.

//  Packages that are created by you or other Go developers
//  live inside the GOPATH,

//  Had you named the package something other than main, like hello for instance,
//  you’d have been telling the compiler this is just a package, not a command.

// When you import code from the standard library, you only need to reference
//  the name of the package, unlike when you import code from outside of the standard
//  library. The compiler will always look for the packages you import at the locations ref
// erenced by the GOROOT and GOPATH environment variables.

// If your main function doesn’t exist in package main,
// the build tools won’t produce an executable.

//  Every code file in Go belongs to a package

//  packages define a unit of compiled code
//  and their names help provide a level of indirection to the identifiers that
//  are declared inside of them, just like a namespace

//  The package name main has special meaning in Go. It designates to the Go command
//  that this package is intended to be compiled into a binary executable. All of the exe
// cutable programs you build in Go must have a package called main.

//  When the main package is encountered by the compiler, it must also find a function
//  called main(); otherwise a binary executable won’t be created. The main() function is
//  the entry point for the program, so without one, the program has no starting point. The
//  name of the final binary will take the name of the directory the main package is
//  declared in.

//  Remember that in Go, a command is any executable program, in
//  contrast to a package, which generally means an importable semantic unit of
//  functionality.

//  In Go if something starts with a capital letter that means other packages (and programs) are able to see it.
//  If we had named the function average instead of Average our main program would not have been able to see it.

// go modules are collection of packages
//  Go Modules are needed to create custom packages.
// the question is why?? why do we need go modules to create custom go packages ??
// The answer is the import path for the custom package we create is derived from the name of the go module
//  In addition to this, all the other third-party packages(such as source code from github) along with their versions
//   which our application uses will be managed by the go.mod file.
// This go.mod file is created when we create a new module

import (
	"fmt"
	"level_0/packages"
)

func init() {
	// All init functions in any code file that are part of the program will
	//  get called before the main function
	fmt.Println("before main")

}

func main() {
	var name string

	// Ask for the user's name
	fmt.Print("Enter your name: ")
	// Read the name input from the user
	fmt.Scanln(&name)
	packages.Hello()

	// Print a greeting message
	fmt.Printf("Hello, %s! Welcome to the world of Go.\n", name)
}
