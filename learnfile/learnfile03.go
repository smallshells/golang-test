package learnfile

import (
	"fmt"
	"io/ioutil"
)

//------------ 利用ioutil读取整个文件 -------------------
/*
	ioutil
 */

func ReadFile03() {
	content, err := ioutil.ReadFile("./learnfile/log")
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	fmt.Println(string(content))
}
