package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print("error=", err)
		return
	}
	defer conn.Close()
	//接收用户输入，写到服务器
	go func() {
		str := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.stdin.err=", err)
				return
			}
			conn.Write(str[:n])
		}
	}()
	//读取服务器内容
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
