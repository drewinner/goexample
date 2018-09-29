package lib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/mikemintang/go-curl"
)

/**
	测试过程发现tcp长连接比较耗资源，改为http请求
	@author:wanghongli
	@since:2018/09/29
 */
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
		//fmt.Println(string(id), string(msg), string(ispush), string(userid), string(create_time))
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
	if _, ok := CluNodeMap[nodeName]; ok == true {

		//组装结构体
		pushMsg := PushMsgStru{colum[0], colum[1], colum[2], colum[3], colum[4]}
		//转换成map
		postData := Struct2Map(pushMsg)
		//请求webservice 通过webservice发送消息
		httpReq := curl.NewRequest()
		//将json转换为map
		httpRep, err := httpReq.SetUrl("http://127.0.0.1:8080").SetPostData(postData).Post()
		if err != nil {
			panic("curl error := " + err.Error())
		} else {
			if httpRep.IsOk() {
				fmt.Println("http url right := ", httpRep.Body)
			} else {
				fmt.Println("http url err :=", httpRep.Raw)
			}
		}
	} else {
		panic("search node err and the nodeName is " + nodeName)
	}

	defer c.Close()

	ch <- 0

}
