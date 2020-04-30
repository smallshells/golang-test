package learnfile

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//------------ 利用bufio读取文件 -------------------
/*
	bufio
 */

func ReadFile02() {
	file, err := os.Open("./learnfile/log")
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	//为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。
	defer file.Close()
	//var tmp = make([]byte,128)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}
