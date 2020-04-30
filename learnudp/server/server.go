package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// server 端
func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"), //net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("启动server失败，err：", err)
		return
	}
	defer listener.Close()
	//循环收发数据

	for {
		var buf [1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收消息失败：err:", err)
			return
		}
		fmt.Printf("接收到来自%v的消息:%v\n", addr, string(buf[:n]))
		//回复消息
		var input string
		//fmt.Scan(&input)
		//fmt.Fprintln(conn,"你好！")
		reader := bufio.NewReader(os.Stdin)  //从标准输入获取输入
		input, err = reader.ReadString('\n') //读取内容直到'\n'
		n, err = listener.WriteToUDP([]byte(input), addr)
		if err != nil {
			fmt.Println("回复消息失败，err:", err)
			return
		}
	}

}
