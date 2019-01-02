package main

import (
	"flag"
	"runtime"
)

var (
	config *Config
	Debug  bool
)

func main() {
	flag.Parse()
	//初始化配置文件
	config = &Config{}
	config.GetConf()
	//设置开发模式
	Debug = config.Base.Debug
	//设置cpu使用核数
	runtime.GOMAXPROCS(config.Base.MaxProcess)
	//

}
