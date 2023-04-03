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
		fileVal := fmt.Sprintf("%s.%d", funcVal, path.Base(entry.Caller.File), entry.Caller.Line)
		// entry.Message是调用log方法时传入函数的日志信息
		fmt.Fprintf(b, "%s [%s] \x1b[0;%dm%s\x1b[0m %s %s %s\n", log.Prefix(), Timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[0;%dm%s\x1b[0m %s \n", Timestamp, levelColor, entry.Level, entry.Message)
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
