package lib

import (
	"fmt"
	"net"
	"os"
)

func Client() {

	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dail err = ", err)
		return
	}
	defer conn.Close()
	//接收用户输入，发送到服务器
	go func() {
		str := make([]byte, 2048)
		n, err := os.Stdin.Read(str)
		if err != nil {
			fmt.Println("os.stdin.err=", err)
			return
		}
		conn.Write(str[:n])
	}()
	//读取服务器内容
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
