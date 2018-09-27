package lib

import (
	"fmt"
	"net"
	"strings"
)

func goHandler(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr)

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//将消息转化为Message对象
		m,err := Jsondecode(string(buf[:n]))
		if err != nil {
			fmt.Println("json error = ",err)
		}

		if m.Msg_type == "cmd" {
			fmt.Println("insert into redis")
			m.Msg_conn = &conn
			onLineUser(m)
		}
		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
func Server() {

	listen, err := net.Listen("tcp", ":8000")
	defer listen.Close()
	if err != nil {
		fmt.Println("the listen is error ", err)
		return
	}

	//接收多个用户
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn err =", err)
			continue
		}
		//处理用户请求
		go goHandler(conn)
	}
}
