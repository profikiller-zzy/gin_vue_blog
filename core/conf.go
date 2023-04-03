package core

import (
	"fmt"
	"gin_vue_blog_AfterEnd/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func InitConfig() *config.Config {
	// 使用ioutil导入配置文件，使用yaml.Unmarshal将配置文件反序列化读取到结构体中
	const ConfigFile = "setting.yaml"
	config := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf file error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, config)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err) // log.Fatalf()用于记录一条严重的错误消息，并且终止程序运行
	}
	fmt.Println("config yamlFile Init success.")
	fmt.Println(config)
	return config
}
