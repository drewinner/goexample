package lib

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckRoomCliError(err error) {
	if err != nil {
		fmt.Printf("error : %s", err.Error())
		os.Exit(1) //0代表成功 非零代表失败
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			conn.Close()
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Printf("client connect err = %s", err)
			break
		}
	}
}
func ChatRoomClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	CheckRoomCliError(err)
	defer conn.Close()

	go MessageSend(conn)

	buf := make([]byte, 1024)

	for {
		length, err := conn.Read(buf)
		CheckRoomCliError(err)
		fmt.Println("receive server message content:", string(buf[:length]))
	}
	fmt.Println("client end !")
}
