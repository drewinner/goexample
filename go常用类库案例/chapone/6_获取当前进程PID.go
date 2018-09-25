package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {

	pid := os.Getpid()
	fmt.Printf("Process PID %d\n", pid)

	if runtime.GOOS != "windows" {
		prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
		out, err := prc.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	} else {
		fmt.Println("the os is not *nix")
	}

}
