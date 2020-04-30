package mylogger

import (
	"fmt"
	"os"
	"time"
)

//往终端打印

type ConsoleLogger struct {
	level Level
}

//文件日志结构体构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	return &ConsoleLogger{
		level: logLevel,
	}
}

//公共记录日志的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) //得到用户要记录的日志
	//日志格式：[时间][文件：行号][函数名][日止级别] 日志信息
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	levelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		timeStr, fileName, line, funcName, levelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg) //利用fmt包将msg字符串写入f.file文件中
}

//Debug方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

//Info方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

//Warning方法
func (c *ConsoleLogger) Warning(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

//Error方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

//Fatal方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

//终端标准输出不需要关闭
func (c *ConsoleLogger) Close() {}
