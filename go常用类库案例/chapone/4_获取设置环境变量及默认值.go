package main

import (
	"fmt"
	"log"
	"os"
)

// DB_CONN=db:/user/db && go run 4_获取设置环境变量及默认值.go
// os.Getenv 如果没有返回空字符串
// os.LookupEnv(key) 返回对应的值 ，及是否设置
func main() {

	test := os.Getenv("CLASSPATH")
	fmt.Println(test)
	connStr := os.Getenv("DB_CONN")
	log.Printf("connection string :%s\n", connStr)

	key := "DB_CONN"
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("the env variable %s is not set .\n", key)
	}
	fmt.Println(connStr)

	key2 := "DB_CONN"
	os.Setenv(key2, "mysql://root:12334@localhsot")
	val := GetEvnDefault(key2, "defaultmysql://localhsot")
	log.Println("The value is :", val)

	os.Unsetenv(key)
	val = GetEvnDefault(key, "mysql://default@localhost")
	log.Println(val)
}

func GetEvnDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
