package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	confFile string
)

type Base struct {
	MaxProcess  int    `yaml:"maxProcess"`
	Debug       bool   `yaml:"debug"`
	ClusterName string `yaml:"clusterName"`
	NodeName    string `yaml:"nodeName"`
}

type BucketOptions struct {
	Num int `yaml:"num"`
	//每个bucket保存的最多长连接数量
	ChannelNum int `yaml:"channelNum"`
}


type Config struct {
	Base          Base          `yaml:"base"`
	WebSocket     []string      `yaml:"webSocket"`
	BucketOptions BucketOptions `yaml:"bucketOptions"`
}

func init() {
	flag.StringVar(&confFile, "c", "./config.yaml", "set the comment config file")
}

/**
  *创建config对象
 */
func (cometConfig *Config) GetConf() *Config {

	yamlFile, err := ioutil.ReadFile(confFile)

	if err != nil {
		panic("err1" + err.Error())
	}
	err = yaml.UnmarshalStrict(yamlFile, cometConfig)
	if err != nil {
		panic("err2" + err.Error())
	}
	return cometConfig

}
