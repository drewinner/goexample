package lib

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
)

/**
	用户和连接map对应关系
	@author:wanghongli
	@since:2018/09/28
*/
type UCMap struct {
	ClusterName string `clusterName:"string"` //集群名称
	NodeName    string `nodeName:"string"`    //node节点名称
	ConnKey     string `connKey:"string"`     //连接key--token
	MsgType     string `msgType:"string"`     //消息类型cmd为指令消息,text为文本消息
	MsgBody     string `MsgBody:"string"`     //消息
}

type PushMsgStru struct {
	Id         string `id:"string"`
	Msg        string `id:"string"`
	IsPush     string `id:"string"`
	UserId     string `userId:"string"`
	CreateTime string `createTime:"string"`
}

/**
	将接收到的消息转化为结构体类型
	@author:wanghongli
	@since:2018/09/27
 */
func Jsondecode(jsonMsg string) (message UCMap, err error) {
	jsonDec := json.NewDecoder(strings.NewReader(jsonMsg))
	var m UCMap
	err = jsonDec.Decode(&m)
	if err != nil {
		log.Println("msg is not json", jsonMsg)
		return m, err
	}
	return m, nil

}

/**
	将结构体转化为json，此方法有问题--没有调用
	@todo:有问题，想办法实现接口形式调用 2018/09/28
	@author:wanghongli
	@since:2018/09/27
*/
func Jsonencode(i interface{}) (jsonMsg string, err error) {

	j, err := json.Marshal(i)
	if err != nil {
		log.Println("struct change to json err = ", err)
		return "", err
	}
	return string(j), nil

}

/**
将结构体转换成map
@author:wanghongli
@since:2018/09/29
*/
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
