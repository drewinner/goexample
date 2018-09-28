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
	//连接服务器成功之后，将自己的token发送给服务器端
	msg := `{"ClusterName":"push_cluster","NodeName":"node1","ConnKey":"user1","MsgType":"cmd","MsgBody":"1"}`
	byteToken := []byte(msg)
	conn.Write(byteToken)
	//接收用户输入，发送到服务器
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
