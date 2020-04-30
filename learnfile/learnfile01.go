package learnfile

import (
	"fmt"
	"os"
)

//------------ 文件打开/关闭 读取 -------------------
/*
	os.Open() 函数能够打开一个文件，返回一个 *File 和一个 err 。
	对得到的文件实例调用 close() 方法能够关闭文件。
	file.read()读取文件，read方法定义格式如下：
		func (f *File) Read(b []byte) (n int, err error)
		它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回 0 和 io.EOF 。

*/

func ReadFile01() {
	file, err := os.Open("./learnfile/log")
	if err != nil {
		fmt.Println("Open file failed!,err:", err)
		return
	}
	//为了防止文件忘记关闭，我们通常使用defer注册文件关闭语句。
	defer file.Close()
	//var tmp = make([]byte,128)
	var tmp [128]byte
	for {
		n, err1 := file.Read(tmp[:])
		if err1 != nil {
			fmt.Println("read file failed!,err:", err1)
			return
		}
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
