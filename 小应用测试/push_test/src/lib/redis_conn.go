package lib

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func ConnRedis() (conn redis.Conn) {
	os.Setenv("REDIS_URL", "redis://192.168.0.222:6379")

	c, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println("the redis conn err= ", err)
		return
	}
	return c
}