package main

import (
	"fmt"
	"io"
	"net"
)

// http client端
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:9090")
	if err != nil {
		fmt.Println("访问失败")
		return
	}
	defer conn.Close()
	//发送数据
	//conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	//接收数据
	var buf [1024]byte
	for{
		n,err := conn.Read(buf[:])
		if err == io.EOF{
			fmt.Print(string(buf[:n]))
			return
		}
		if err != nil {
			fmt.Println("接收数据失败，err",err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
