package main

import (
	"fmt"
)

func main() {
	num := 1

	pointer := &num //Points to num
	fmt.Println(*pointer)
}
