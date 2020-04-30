package mylog

import (
	"path"
	"runtime"
)

func getCallerInfo() (funcName string, fileName string, line int) {
	pc, fileName, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName)
	fileName = path.Base(fileName)
	return
}
