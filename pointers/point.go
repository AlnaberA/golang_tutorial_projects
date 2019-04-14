package main

import (
	"fmt"
)

func main() {
	num1 := 0
	num2 := 100

	pointer := &num1 //Points to num1
	fmt.Println(*pointer)
	fmt.Println(pointer)

	*pointer = 1 //Set value of pointer to 2 through the pointer.
	fmt.Println(*pointer)
	fmt.Println(pointer)

	pointer = &num2
	fmt.Println(*pointer)
	fmt.Println(pointer)
}
