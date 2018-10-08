package main

import "os"
import "lib"

func main() {
	arg := os.Args[1]
	if arg == "s" {
		lib.ChatRoomServer()
	} else if arg == "c" {
		lib.ChatRoomClient()
	}
}
