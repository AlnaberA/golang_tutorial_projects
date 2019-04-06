package main

import "os/exec"

func main() {
	// Command function accepts args
	cmd := exec.Command("ls", "-a")

	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		// If there is an error the program exits the program with status 1.
	}

	print(string(stdout))
}
