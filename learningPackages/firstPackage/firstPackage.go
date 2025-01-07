package firstPackage

// first we created the folder learningPackages
// then,did $ go mod init learningPackages
// so that means learningPackgaes is our module
// now after that we created the folder firstPackage
// and then this file
// to use third party packages,you have to have module initialized
// to initialize a module,we use go mod init and pass a fully qualified
// name to it
// for example,if we wanted to host module on github under username sidhartGautam
// we can initialize our module like this
// $ go mod init github.com/sidhartGautam/myModule

// to use a third party module,let say module uuid
// we do this --> go get github.com/google/uuid
// this will place this moudle or package at directory $GOPATH/pkg/mod/
// You can run go env GOMODCACHE to see where Go will put downloaded modules.

// The go get command noticed the go.mod file in your current directory and updated
// it to reflect your program’s new dependency, including its version. Now you can
// use this package in your package.

// we make this function inside firstpackage and first letter is capital
// meaning it can be used by other packages

// it is a common practise that entry point of a package is the packageName.go
// file

// one important things is,diffrent directory in a same module can have files
// with same package name,but it does not mean that they same packages
// in go,package is determined by the directory location and not jyt by name,
// so even if two directory say same package they are diffrent as they
// have diffrent location
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
