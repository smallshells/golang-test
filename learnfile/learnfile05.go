package learnfile

import (
	"bufio"
	"fmt"
	"os"
)

//----------------------- 利用bufio.NewWriter 文件写入文件 -------------------
/*
	把缓存中的内容写入
*/
func WriteFile05() {
	file, err := os.OpenFile("./log", os.O_APPEND|os.O_CREATE, 02)
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++{
		writer.WriteString("Hello World!\n")  //先将数据写入缓存
	}
	writer.Flush()  // 将缓存中的内容写入文件
}
