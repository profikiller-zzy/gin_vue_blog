package core

import (
	"fmt"
	"gin_vue_blog_AfterEnd/config"
	"gin_vue_blog_AfterEnd/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const ConfigFile = "setting.yaml"

func InitConfig() *config.Config {
	// 使用ioutil导入配置文件，使用yaml.Unmarshal将配置文件反序列化读取到结构体中
	config := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf file error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, config)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err) // log.Fatalf()用于记录一条严重的错误消息，并且终止程序运行
	}
	//fmt.Println("config yamlFile Init success.")
	return config
}

func SetYaml() error {

	// 将结构体编码为yaml格式
	siteInfoYaml, err := yaml.Marshal(&global.Config)
	if err != nil {
		return err
	}

	// 将yaml内容写入文件
	err = ioutil.WriteFile(ConfigFile, siteInfoYaml, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}
