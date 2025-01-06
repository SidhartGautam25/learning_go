package main

import (
	"fmt"
	"learningPackages/firstPackage"
	"level_0/packages"
)

// Main is our second package actually
// in firstpackage,we have stated that entryfile should be named of package or
// directory name
// and in second package,we are not following that rule
// the reason is this is our executable package or program
// so here,instead of Main,we can name it as secondPackage and leave the
// name of file as main.go and can say this package as package main

// now in this second package or main package,we will try to learn two
// basic things.
// first is,how to use the package of the same module,in this case that
// is firstPackage
// and second is how to use package of some diffrent module,in this case
// of module level0

func main() {
	var p = 4.0
	var r = 2.0
	var t = 3.0
	//here we are using local package firstPackage
	var result = firstPackage.Calculate(p, r, t)
	fmt.Println(result)

	// now we have to use other module package
	// in this case we need to use hello fucntion of packages package of
	// level0 module
	// for that,first we need to update the go.mod file
	// there we need to do this -->
	// $ replace level_0 => ../level_0
	// this simply means that level_0 module can be found at ../level_0 location
	// now we need to do $ go get level_0 as to use it we have to bring it to this module
	// now to use packages package of level_0 module we will import level_0/packages

	packages.Hello()

}
