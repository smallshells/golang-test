package learnfile

import (
	"fmt"
	"os"
)

//----------------------- 文件中间插入内容 -------------------
/*
	新建一个
*/
func WriteFile07() {
	file, err := os.OpenFile("./log",os.O_RDWR,0644)
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	defer file.Close()
	file.Seek(2, 0)
	file.Write([]byte("hello"))  //会覆盖原来的值
	var ret [4]byte
	n, err1 := file.Read(ret[:])
	if err1 != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	fmt.Println(string(ret[:n]))
}


func WriteFile08(){
	//打开要插入的主文件
	file, err := os.OpenFile("./log",os.O_RDWR,0644)
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	defer file.Close()
	//因为没有办法直接在文件中间直接插入内容，所以要借助临时文件
	tmpFile, err := os.OpenFile("./sb.tmp",os.O_CREATE | os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("Create tmp file failed!,err:", err)
		return
	}
	defer tmpFile.Close()
	//读源文件内容读出
	var ret [4]byte
	n, err1 := file.Read(ret[:])
	if err1 != nil {
		fmt.Printf("read from file failed, err:%v", err)
		return
	}
	//写入临时文件
	tmpFile.Write(ret[:n])
	//再写入要插入的内容
	tmpFile.Write([]byte("hello world"))
	//紧接着把源文件后续内容写入临时文件
	for{
		n2, err2 := file.Read(ret[:])
		if err2 != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		if n2 < 4 {
			return
		}
		tmpFile.Write(ret[:n2])
	}
	//把临时文件内容读入到源文件


	//删除临时文件
}