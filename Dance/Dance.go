package main

import (
	"fmt"
	"time"
)

func main() {
	//https://blog.golang.org/gos-declaration-syntax
	//https://tour.golang.org/basics/8
	var amountOfDances int

	fmt.Print("Enter amount of times you want to dance: ")

	//Parameters are passed by value in Go. amountOfDances is the integer value in amountOfDances, &amountOfDances is a pointer to amountOfDances
	fmt.Scan(&amountOfDances)

	dance(amountOfDances)
}

func dance(amountOfDances int) {
	var frames [2]string

	frames[0] = "♪┏(・o･)┛♪"
	frames[1] = "♪┗(･o･)┓♪"

	/*
		Go will infer the type of initialized variables.
		Variables declared without a corresponding initialization are zero-valued.
	*/

	//Switch between frame 0 and 1
	//Here the type will be inferred.
	var dance = false
	/*
		Short-hand notation:
		Ln 25-26 are equivalent
		var dance = false
		dance := false
	*/

	// Amount of dances
	for i := 0; i < amountOfDances; i++ {
		//Logic of selecting the frame.
		if dance == true {
			// Control character used to reset devices position to the beginning of a line. (\r)
			fmt.Printf("\r" + frames[0])
			dance = false
		} else {
			fmt.Printf("\r" + frames[1])
			dance = true
		}
		time.Sleep(1 * time.Second) //Delay
	}
}
