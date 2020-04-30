package main

import (
	"fmt"
	"net"
)

//单独处理连接的函数
func process(conn net.Conn){
	defer conn.Close()
	//从连接中接收数据
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("接收客户端发来的消息失败了，err:",err)
		return
	}
	fmt.Println(">>>",string(buf[:n]))
}

func main() {
	//1.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen : 20000 failed, err: ", err)
		return
	}
	defer listener.Close()
	//2.接收客户端请求建立连接
	for {
		conn, err := listener.Accept()  //如果没有客户端连接就阻塞，一直在等待
		if err != nil {
			fmt.Println("connect failed, err : ", err)
			continue
		}
		//3.创建goroutine处理连接
		go process(conn)
	}
}
