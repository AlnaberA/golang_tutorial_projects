package main

import "os/exec"

func main() {
	// Command function accepts args
	cmd := exec.Command("ls", "-a")

	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))
}
