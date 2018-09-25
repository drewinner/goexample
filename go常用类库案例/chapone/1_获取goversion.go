package main

import (
	"fmt"
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO:%s`

func GetGoVersion() {
	log.Printf(info, "Example", runtime.Version())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOROOT())
}
