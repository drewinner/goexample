package lib

import (
	"fmt"
	"net"
	"strings"
)

//map用来存放用户和连接的对应关系--注意限流，map最大个数限制流量
var MsgMap map[string]net.Conn = make(map[string]net.Conn)


func goHandler(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr)

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err111111= ", err)
			return
		}
		//将消息转化为Message对象
		m, err := Jsondecode(string(buf[:n]))
		if err != nil {
			fmt.Println("json error = ", err)
		}

		if m.MsgType == "cmd" && m.MsgBody == "1" {
			fmt.Println("insert into redis")
			onLineUser(m,string(buf[:n]))
			//保存map 用户token和用户conn映射关系
			MsgMap[string(m.ConnKey)] = conn

			fmt.Println(MsgMap)
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
	//开启http服务
	go WebServer()
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
