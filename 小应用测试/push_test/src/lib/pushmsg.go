package lib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net"
)

var ch chan int = make(chan int)
var CluNodeMap map[string]string = make(map[string]string)

func init() {
	CluNodeMap["node1"] = ":8000"
}

/**
	发送消息，暂时用mysql-->后期改为kafka->消息分发
	@author:wanghongli
	@since:2018/09/27
 */
func PushMsg() {
	fmt.Println(CluNodeMap)
	//获取数据库连接
	db := ConnMysql()
	//查询即将发送的消息
	defer db.Close()

	rows, err := db.Query("select * from pre_msg where ispush=0 order by id desc limit 10")
	if err != nil {
		panic(err.Error())
	}

	var id, msg, ispush, userid, create_time []byte
	for rows.Next() {
		err = rows.Scan(&id, &msg, &ispush, &userid, &create_time)
		if err != nil {
			panic("read mysql error = " + err.Error())
		}
		fmt.Println(string(id), string(msg), string(ispush), string(userid), string(create_time))
		go handlePush(string(id), string(msg), string(ispush), string(userid), string(create_time))
		<-ch
	}
}

/**
	新开启线程，发送消息
	@todo:将来做消息分发用
	@author:wanghongli
	@since:2018/09/27
*/
func handlePush(colum ...string) {

	//根据userid获取对应的tcp conn
	c := ConnRedis()
	//val 里保存的是集群信息，节点信息
	val, err := redis.String(c.Do("get", colum[3]))
	if err != nil {
		panic("push err happen = " + err.Error())
	}
	//根据集群和节点信息分发到对应的节点
	jsonVal, err := Jsondecode(val)
	if err != nil {
		panic("json err = " + err.Error())
	}

	//clusterName := jsonVal.ClusterName //暂时不用
	nodeName := jsonVal.NodeName
	//根据nodeName获取对应的ip和端口号
	if nodeNameV, ok := CluNodeMap[nodeName]; ok == true {

		connn, err := net.Dial("tcp", nodeNameV)
		if err != nil {
			panic("pushmsg error " + err.Error())
		}
		//组装结构体
		pushMsg := PushMsgStru{colum[0], colum[1], colum[2], colum[3], colum[4]}
		//转换成json
		jsonPushMsg, err := Jsonencode(pushMsg)
		fmt.Println(jsonPushMsg)
		n,err :=connn.Write([]byte(jsonPushMsg))
		if err != nil {
			fmt.Println("push error !")
		}
		fmt.Println(n)
		//defer connn.Close()
	} else {
		panic("search node err and the nodeName is " + nodeName)
	}

	defer c.Close()
	ch <- 0

}
