package main

//Package exec runs external commands.
import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("date")
	//Output() returns two values stdout and err so two varibles need to be assigned to the method.
	output, err := cmd.Output()
	if err != nil {
		print(err.Error())
	}
	//%s string format
	fmt.Printf("The date is %s\n", output)
}
