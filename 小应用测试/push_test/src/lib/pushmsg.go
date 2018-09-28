package lib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net"
)

var ch chan int = make(chan int)
var cluNodeMap map[string]string



/**
	发送消息，暂时用mysql-->后期改为kafka->消息分发
	@author:wanghongli
	@since:2018/09/27
 */
func PushMsg() {
	fmt.Println(cluNodeMap)
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
			panic(err.Error())
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
	//和主机建立链接
	conn,err := net.Dial("tcp",":8000")
	if err != nil {
		panic("push err happen = "+err.Error())
	}
	//根据集群和节点信息分发到对应的节点

	conn.Write([]byte(val))

	defer conn.Close()
	defer c.Close()
	//if err != nil {
	//	//将来记录到其它存储中
	//	log.Println("redis get key err = ", err)
	//} else {
	//	//根据val获取conn资源
	//	fmt.Println("-------------")
	//	fmt.Println(MsgMap)
	//	if mapV, ok := MsgMap[colum[3]]; ok == true {
	//		fmt.Println("push start ....")
	//		mapV.Write([]byte("开始推送啦！"))
	//	} else {
	//		fmt.Println("推送失败")
	//	}
	//}
	ch <- 0

}
