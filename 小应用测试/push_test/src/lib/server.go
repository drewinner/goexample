package lib

import (
	"fmt"
	"net"
)
func goHandler(conn net.Conn){
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	var b []byte
	b = []byte{'a','b'}
	conn.Write(b)


}
func Server(){

	listen,err := net.Listen("tcp",":8000")
	defer listen.Close()
	if err != nil {
		fmt.Println("the listen is error ",err)
		return
	}
	defer listen.Close()

	for {
		conn,err := listen.Accept()
		if err !=nil {
			fmt.Println("conn err =",err)
			continue
		}
		goHandler(conn)
	}
}
