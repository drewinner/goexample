package lib

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

/**
	连接redis
	@author:wanghongli
	@since:2018/09/27
 */
func ConnRedis() (conn redis.Conn) {
	//os.Setenv("REDIS_URL", "redis://192.168.0.222:6379")
	os.Setenv("REDIS_URL", "redis://10.115.88.111:6379")

	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println("the redis conn err= ", err)
		return
	}
	return c
}

/**
	将用户连接，和token以键值对的形式存放到redis中
	@author:wanghongli
	@since:2018/09/27
*/
func onLineUser(m UCMap,normalMsg string) bool {
	c := ConnRedis()
	c.Do("set",m.ConnKey,normalMsg)
	return true
}
