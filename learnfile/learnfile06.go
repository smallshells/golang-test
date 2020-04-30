package learnfile

import (
	"fmt"
	"io/ioutil"
)

//----------------------- 利用 ioutil.WriteFile 文件写入文件 -------------------
/*
	ioutil.WriteFile
*/
func WriteFile06() {
	err := ioutil.WriteFile("./learnfile/log", []byte("I want to go home\n"), 0666)
	if err != nil {
		fmt.Println("Write file failed!, err:", err)
		return
	}
}
