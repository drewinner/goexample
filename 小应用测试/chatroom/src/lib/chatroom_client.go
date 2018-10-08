package lib

import (
	"fmt"
	"net"
	"os"
)
func CheckRoomCliError(err error) {
	if err != nil {
		fmt.Printf("error : %s", err.Error())
		os.Exit(1) //0代表成功 非零代表失败
	}
}
func ChatRoomClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	CheckRoomCliError(err)
	defer conn.Close()

	conn.Write([]byte("hello world"))
	fmt.Println("has send the message!")
}
