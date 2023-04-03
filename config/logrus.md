# logrus自定义日志输出格式  
## 自定义日志输出格式  
在logrus中，使用如下方法设置日志格式  
```
func SetFormatter(formatter Formatter) 
```

实现自定义日志格式，本质上就是实现Formatter接口，然后通过SetFormatter方式将其告知logrus。  
```
type formatter interface {
    Formate(*logrus.Entry) ([]byte, error)
}
```
接口的返回值是byte[], error，即为输出串，关键要搞懂输入参数Entry：  
Entry参数： 
```azure
type Entry struct {
	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	Level Level

	//Calling method, with package name
	Caller *runtime.Frame

	//Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	//When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer
}
```
说明：
 - Data中是key/value形式的数据，是使用WithField设置的日志。
 - Caller是日志调用者相关的信息，可以利用其输出文件名，行号等信息，感兴趣可以参看[《logrus中输出文件名、行号及函数名》](https://blog.csdn.net/qmhball/article/details/116656368)  
例子：  
```azure
type MyFormatter struct {

}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error){
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func main(){
	logrus.SetFormatter(&MyFormatter{})
    logrus.WithField("name", "ball").WithField("say", "hi").Info("info log")
}

//输出
[2021-05-10 17:26:06] [info] info log
```