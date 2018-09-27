package main

import (
	"lib"
	"os"
)

func main() {
	args := os.Args
	if args[1] == "c" {
		lib.Client()
	} else if args[1] == "s" {
		lib.Server()
	} else if args[1] == "push" {
		lib.PushMsg()
	}
}
