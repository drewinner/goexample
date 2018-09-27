package lib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var ch chan int = make(chan int)
/**
	发送消息，暂时用mysql-->后期改为kafka
	@author:wanghongli
	@since:2018/09/27
 */
func PushMsg() {

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
	@author:wanghongli
	@since:2018/09/27
*/
func handlePush(colum ...string) {

	//根据userid获取对应的tcp conn
	c := ConnRedis()
	val, err := redis.String(c.Do("get", colum[3]))
	if err != nil {
		//将来记录到其它存储中
		log.Println("redis get key err = ", err)

	} else {
		//根据val获取conn资源
		fmt.Println("-------------")
		fmt.Println(val)
	}
	ch <- 0

}
