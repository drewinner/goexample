package main

import (
	"fmt"
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO:%s`

func main() {
	log.Printf(info, "Example", runtime.Version())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOOS)
}
