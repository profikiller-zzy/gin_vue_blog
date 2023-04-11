#项目主体结构

![img.png](img.png)

# 基本配置
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
```
type Config struct {
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
```
package core

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

同时我们还需要将配置文件存入global中，因为配置文件应当是全局的，需要现在global中创建相应的结构体：

```package global

import "gin_vue_blog_AfterEnd/config"

var (
	Config *config.Config
)
```

保存的操作已经涵盖在`func InitConfig()`中了  
## 读取配置文件  
core/conf.go:
```azure
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
	//fmt.Println("config yamlFile Init success.")
	return config
}

```

## gorm配置  
core/gorm.go
```azure
package core

import (
	"gin_vue_blog_AfterEnd/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// InitGorm gorm连接到mysql数据库
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql数据库，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	// 设置mysql日志
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) //
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf("[%s] mysql连接失败", dsn)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              // 连接池最大容量
	sqlDb.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}

```

## 日志配置  
由于日志用到了logrus，logrus允许我们自定义日志，包括日志输出格式、日志输出位置、日志输出等级等等。一般步骤为先调用logrus库中的Formatter接口来定义自己的日志格式，然后调用logger的SetFormatter()方法来将日志格式设置为自定义格式：  
```azure
logger.SetFormatter(&MyFormatter{})
```
core/logrus.go:
```azure
package core

import (
	"bytes"
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry logrus.Entry) ([]byte, error)接口方法
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义输出日期格式
	Timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义打印调用日志的函数和行号
		funcVal := entry.Caller.Function
		// entry.Caller.File是调用函数的名称，entry.Caller.Line是调用时调用代码所在的行号
		fileVal := fmt.Sprintf("%s.%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// entry.Message是调用log方法时传入函数的日志信息
		fmt.Fprintf(b, "%s [%s] \x1b[%dm%s\x1b[0m %s %s %s\n", log.Prefix(), Timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm%s\x1b[0m %s\n", log.Prefix(), Timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// InitLogger 返回初始化日志实例
func InitLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stderr)                           // 设置日志内容的输出方式
	log.SetReportCaller(global.Config.Logger.ShowLine) // 设置是否输出调用函数的名称和代码行号
	log.SetFormatter(&LogFormatter{})                  // 设置自己定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil { // 如果设置文件中没有设置日志等级，就使用默认日志等级，默认日志等级为info
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
	InitDefaultLogger() // 配置全局log
	return log
}

func InitDefaultLogger() {
	log := logrus.New()
	log.SetOutput(os.Stderr)                           // 设置日志内容的输出方式
	log.SetReportCaller(global.Config.Logger.ShowLine) // 设置是否输出调用函数的名称和代码行号
	log.SetFormatter(&LogFormatter{})                  // 设置自己定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil { // 如果设置文件中没有设置日志等级，就使用默认日志等级，默认日志等级为info
		level = logrus.InfoLevel
	}
	log.SetLevel(level)
}

```

## 路由配置
 - 路由引擎初始化
```azure
package router

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
    
    // 路由分组
	apiRouter := router.Group("/api/")

	apiRouterGroupApp := RouterGroup{
		RouterGroup: apiRouter,
	}
	// 路由分层
	apiRouterGroupApp.SettingRouter()

	return router
}

```
值得一提的是，上面的代码还体现了路由分层和路由分组。  
在Go中，路由分层和路由分组是两个不同的概念，它们的区别如下：

路由分层（Routing Hierarchies）：指将不同的路由映射到不同的处理器函数或控制器上，以便更好地组织代码。在路由分层中，每个路由都对应一个处理器函数或控制器，用于处理该路由的请求。路由分层可以使代码更易于维护和扩展。

路由分组（Routing Groups）：指将相关的路由分组在一起，以便在它们上面执行共同的操作，例如添加中间件或共享路由参数。在路由分组中，可以将多个路由定义在同一个路由组中，并在路由组上设置中间件或路由参数，以便这些路由共享相同的中间件或参数。路由分组可以使代码更具可读性和可维护性。

简而言之，路由分层是将路由映射到不同的处理器函数或控制器上，以实现代码的组织和维护；而路由分组是将相关的路由分组在一起，以便在它们上面执行共同的操作。两者都可以使代码更易于维护和扩展。  
 - 目前具体的路由文件  
router/setting_info.go
```azure
package router

import (
	"gin_vue_blog_AfterEnd/api"
)

// SettingRouter 系统配置api
func (r RouterGroup) SettingRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	r.GET("/setting/", settingApi.SettingInfoView)
}

```
 - 封装统一的路由响应
```azure
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装了一些gin公共的响应
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := CodeMessage[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(int(code), map[string]interface{}{}, "未知错误", c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, c)
}

```
以上几个函数分别是几个通用的请求响应成功和请求响应失败时向用户返回的response，将他们用函数封装起来便于调用。  
 - 错误码封装
```azure
package response

type ErrorCode int

const (
	SettingsError = ErrorCode(1001)
)

var CodeMessage = map[ErrorCode]string{
	SettingsError: "系统错误",
}

```
也可以将错误码写入json文件中，在程序开始时将json文件中保存的错误码及对应的错误码信息通过json.Unmarshall()读取到定义的map结构中，这段测试代码反应了这种方法：  
```azure
package main

import (
	"encoding/json"
	"fmt"
	"gin_vue_blog_AfterEnd/model/response"
	"io/ioutil"
)

type CodeMsg map[int]string

const FilePath = "model/response/ErrCode.json"

func main() {
	jsonFile, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println(err)
	}

	var codeMsg CodeMsg
	err = json.Unmarshal(jsonFile, &codeMsg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(codeMsg)
	fmt.Println(response.SettingsError)
}

```

# 搭建表结构
下面是整个项目的ER模型图：  
![gvb-ER模型.png](gvb-ER模型.png)

