package util

//To use this package among all of your go applications run the "go install" command.

//Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

//The lower case function name makes the subtract function not exported.
//The upper case Add func is makes the add function visible or exported.
func subtract(a, b int) int {
	return a - b
}
