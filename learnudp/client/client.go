package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("udp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("连接server失败，err:", err)
		return
	}
	defer conn.Close()

	//循环收发数据
	for {
		//发消息
		var input string
		//fmt.Scan(&input)
		//fmt.Fprintln(conn,"你好！")
		reader := bufio.NewReader(os.Stdin)  //从标准输入获取输入
		input, err = reader.ReadString('\n') //读取内容直到'\n'
		n, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送消息失败，err:", err)
			return
		}
		//收消息
		var buf [1024]byte
		n, err = conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取消息失败，err:", err)
			return
		}
		fmt.Println("收到回复：", string(buf[:n]))
	}
}
