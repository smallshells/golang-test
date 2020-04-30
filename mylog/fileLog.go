package mylog

import (
	"fmt"
	"os"
	"time"
)

//FileLogger 往文件里记录日志的结构体
type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

//NewFileLogger 是一个生成文件日志结构体实例的构造函数
func NewFileLogger(lever int, logFilePath, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       lever,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
	flObj.initFileLogger()
	return flObj
}

//专门用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger() {
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filepath))
	}
	f.logFile = file
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > DEBUG {
		return
	}
	funcName, fileName, line := getCallerInfo()
	//日志的格式要丰富起来：
	//时间 日志级别 哪个文件哪一行哪一个函数 日志信息
	//[2020-4-15 22:19:25] [DEBUG] main.go [14] id为10的用户一直在尝试登陆
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	format = fmt.Sprintf("[%s] [%s] [%s:%s] [%d] %s",
		nowStr, getLevelStr(f.level), funcName, fileName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile)
}

func (f *FileLogger) Info(msg string) {
	f.logFile.WriteString(msg)
}

func (f *FileLogger) Error(msg string) {
	f.logFile.WriteString(msg)
}
