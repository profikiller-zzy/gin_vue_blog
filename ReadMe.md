#项目主体结构

![img.png](img.png)

## 配置文件的编写和读取

- 配置文件编写：
```mysql:
  host: 127.0.0.1
  port: 3306
  db: gin_vue_blog
  user: root
  password: root
  log_level: dev
logger:
  level: info
  prefix: '[gin_vue_blog]'
  director: log
  show-line: ture
  log-in-console: true
system:
  host: "0.0.0.0"
  port: 8080
  env: dev
  ```

 - 对应结构体
```type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	//Logger Logger `yaml:"logger"`
	//System System `yaml:"system"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级是指日志消息的重要性和优先性
}
```
 - 读取配置文件  
   使用ioutil导入配置文件，使用yaml.Unmarshal将配置文件反序列化读取到结构体中
```package core

import (
	"fmt"
	"gin_vue_blog_AfterEnd/config"
	"gin_vue_blog_AfterEnd/global"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func InitConfig() {
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
	global.Config = config // 把读取到的配置文件存放到global中，配置文件应当是全局的
}
```

同时我们还需要将配置文件存入global中，应为配置文件应当是全局的，需要现在global中创建相应的结构体：

```package global

import "gin_vue_blog_AfterEnd/config"

var (
	Config *config.Config
)
```

保存的操作已经涵盖在`func InitConfig()`中了