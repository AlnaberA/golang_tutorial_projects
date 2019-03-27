//package <packagename>
//A package is a directory inside your Go workspace containing one or more Go source files, or other Go packages.
//Programs start running in package main
package main

//Package fmt(format) implements formatted I/O with functions comparable to C's printf and scanf.
import (
	"fmt"
)

// entry point for the application
func main() {
	fmt.Println("Hello World.")
}
