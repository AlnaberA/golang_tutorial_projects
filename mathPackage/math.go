package mathPackage

/*
To use this package among all of your go applications run the "go install" command.
"go install" will not work outside of the Go workspace.
When the "go install" command runs an *.a file is generated in the go/pkg dir
*/

//Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

//The lower case function name makes the subtract function not exported.
//The upper case Add func is makes the add function visible or exported.
func subtract(a, b int) int {
	return a - b
}
