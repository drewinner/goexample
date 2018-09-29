package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

/**
	web服务器--简单版本
	@author:wanghongli
	@since:2018/09/29
 */
func WebServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	jsonString := string(body)
	//转换为json对象
	var pushMsgStru PushMsgStru
	jsonDec := json.NewDecoder(strings.NewReader(jsonString))
	err := jsonDec.Decode(&pushMsgStru)
	if err != nil {
		w.Write([]byte("error"))
	}
	//找到对应连接，推送消息
	token := pushMsgStru.UserId
	msg := pushMsgStru.Msg
	if conn, ok := MsgMap[token]; ok == true {
		conn.Write([]byte(msg))
	} else {
		fmt.Println("push err ! = ")
	}
}
