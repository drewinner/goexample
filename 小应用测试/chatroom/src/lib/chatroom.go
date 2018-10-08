package lib

import (
	"fmt"
	"net"
	"os"
)

func CheckRoomError(err error) {
	if err != nil {
		fmt.Printf("error : %s", err.Error())
		os.Exit(1) //0代表成功 非零代表失败
	}
}

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		CheckRoomError(err)
		if numOfBytes != 0 {
			fmt.Printf("Has received this message : %s", string(buf[:numOfBytes]))
		}
	}

}
func ChatRoomServer() {
	listenSocket, err := net.Listen("tcp", "127.0.0.1:8000")
	CheckRoomError(err)
	defer listenSocket.Close()

	for {
		conn, err := listenSocket.Accept()
		if err != nil {
			continue
		}
		go ProcessInfo(conn)
	}
}
