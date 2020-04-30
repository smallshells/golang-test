package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//这是往文件里面写日志的代码

// FileLogger 文件日志结构体
type FileLogger struct {
	level    Level
	filePath string
	fileName string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

//文件日志结构体构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxSize int64, ) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fl := &FileLogger{
		level:    logLevel,
		filePath: filePath,
		fileName: fileName,
		maxSize:  maxSize,
	}
	fl.initFile() //
	return fl
}

//将制定的日志文件打开，赋值给结构体
func (f *FileLogger) initFile() {
	//打开日志文件
	logFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("open logfile %s failed, %v", logFileName, err))
	}
	f.file = fileObj
	//打开错误日志文件
	errLogFileName := fmt.Sprintf("%s.err", logFileName)
	errFileObj, err := os.OpenFile(errLogFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("open logfile %s failed, %v", logFileName, err))
	}
	f.errFile = errFileObj
}

//封装一个切分日志文件的方法
func (f *FileLogger) splitLogFile(file *os.File) {
	FileInfo, _ := file.Stat()
	fileSize := FileInfo.Size()
	if fileSize >= f.maxSize {
		//切分文件
		fileName := file.Name() //拿到文件的完整路径
		backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
		//1.把原来的文件关闭
		file.Close()
		//2.备份原来的文件
		os.Rename(fileName, backupName)
		//3.新建一个文件
		fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
		if err != nil {
			panic(fmt.Errorf("open logfile %s failed, %v", fileName, err))
		}
		file = fileObj
	}
}

//公共记录日志的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...) //得到用户要记录的日志
	//日志格式：[时间][文件：行号][函数名][日止级别] 日志信息
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	levelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		timeStr, fileName, line, funcName, levelStr, msg)
	//往文件里写之前要做一个检查，检查当前日志文件的大小是否超过了maxSize
	f.splitLogFile(f.file)
	fmt.Fprintln(f.file, logMsg) //利用fmt包将msg字符串写入f.file文件中
	//如果是error或者fatal级别的日志还要记录到 f.errFile
	if level >= ErrorLevel {
		f.splitLogFile(f.errFile)
		fmt.Fprintln(f.errFile, logMsg)
	}
}

//Debug方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

//Info方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

//Warning方法
func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

//Error方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

//Fatal方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

// Close 关闭日志文件句柄
func (f *FileLogger) Close(){
	f.file.Close()
	f.errFile.Close()
}