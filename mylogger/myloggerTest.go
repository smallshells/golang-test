package mylogger

var logger Logger   //定义一个接口，日志不管用文件还是终端，都不用改后面的代码

func TestMyLogger(){
	logger = NewFileLogger("debug","./mylogger","test.log",1000)
	defer logger.Close()
	//logger = NewConsoleLogger("debug")
	logger.Debug("我是debug")
	logger.Info("我是info")
	logger.Error("我是error")
}