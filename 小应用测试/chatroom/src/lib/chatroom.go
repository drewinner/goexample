package lib

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var onlineConns = make(map[string]net.Conn)
//消息队列
var messageQueue  = make(chan string, 10000)
var quitChannel  = make(chan bool)

func CheckRoomError(err error) {
	if err != nil {
		fmt.Printf("error : %s", err.Error())
		os.Exit(1) //0代表成功 非零代表失败
	}
}

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer func(conn net.Conn){
		addr := fmt.Sprintf("%s",conn.RemoteAddr())
		delete(onlineConns,addr)
		conn.Close()
	}(conn)

	for {
		numOfBytes, err := conn.Read(buf)
		CheckRoomError(err)
		if numOfBytes != 0 {
			message := string(buf[:numOfBytes])
			messageQueue <- message
		}
	}

}

//消费消息队列中的消息
func ConsumerMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChannel: //退出
			break
		}
	}
}

func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		addr = strings.Trim(addr, " ")
		sendMessage := contents[1]
		sendMessage = strings.Join(contents[1:],"#")
		if conn, ok := onlineConns[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Println("online conns send failer ! ")
			}
		}
	}
}
func ChatRoomServer() {
	listenSocket, err := net.Listen("tcp", "127.0.0.1:8000")
	CheckRoomError(err)
	defer listenSocket.Close()
	fmt.Println("Server is waiting ... ")
	go ConsumerMessage()
	for {
		conn, err := listenSocket.Accept()
		//将conn存储到映射表onlineConns
		remoteAddr := conn.RemoteAddr().String()
		onlineConns[remoteAddr] = conn
		if err != nil {
			continue
		}
		for c := range onlineConns {
			fmt.Println(c)
		}
		go ProcessInfo(conn)
	}
}
