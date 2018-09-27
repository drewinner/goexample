package lib

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

//json code and json decode

type Message struct {
	Msg_conn *net.Conn                   //连接地址
	Msg_type string `msg_type:"string"` //消息类型 cmd特殊为指令类型消息
	Msg_body string `msg_body:"string"` //消息体,如果是cmd，内容为token
}

/**
	将接收到的消息转化为结构体类型
	@author:wanghongli
	@since:2018/09/27
 */
func Jsondecode(jsonMsg string) (message Message, err error) {
	jsonDec := json.NewDecoder(strings.NewReader(jsonMsg))
	var m Message
	err = jsonDec.Decode(&m)
	if err != nil {
		log.Println("msg is not json", jsonMsg)
		return m, err
	}
	return m, nil

}

/**
	将结构体转化为json，此方法有问题--没有调用
	@author:wanghongli
	@since:2018/09/27
*/
func Jsonencode(m Message) (jsonMsg string, err error) {

	j, err := json.Marshal(m)
	if err != nil {
		log.Println("struct change to json err = ", err)
		return "", err
	}
	return string(j), nil

}
