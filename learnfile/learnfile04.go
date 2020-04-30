package learnfile

import (
	"fmt"
	"os"
)

//----------------------- 利用OpenFile 文件写入文件 -------------------
/*
	func OpenFile(name string, flag int, perm fileMode)(*File, error){
		……
	}
	name:要打开的文件名；flag：打开文件的模式。有以下几种：
		os.O_WRONLY	    只写
		os.O_CREATE     创建文件
		os.O_RDONLY     只读
		os.O_RDWR       读写
		os.O_TRUNC      清空
		0s.O_APPEND     追加
	perm:文件权限，一个八进制数。r(读)04, w(写)02, x(执行)01
*/
func WriteFile04() {
	file, err := os.OpenFile("./learnfile/log", os.O_APPEND|os.O_CREATE, 02)
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	defer file.Close()
	//write
	file.Write([]byte("I want to go shopping\n"))
	//writeString
	file.WriteString("My favorite food is rice!\n")
}
