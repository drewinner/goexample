package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//处理用户请求
func goHandler(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println("remote connect successful")
	//接收用户输入，写到客户端
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
	//读取用户数据
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}

		fmt.Printf("[%s] send and the read buf = %s\n ", addr, string(buf[:n]))
		if "exit" == string(buf[:n-1]) {
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		goHandler(conn)
	}
}
