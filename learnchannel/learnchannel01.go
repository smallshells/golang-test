package learnchannel

import "fmt"

//------------------------------------- channel -------------------------------------
/*
channel:
	单纯地将函数并发执行没有意义，函数与函数间需要交换数据才能体现并发执行函数的意义。
	虽然可以使用共享内存进行数据交换，但是共享内存在不同的 goroutine 中容易发生竞态问题。
	为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。
	Go语言的并发模型是CSP，提倡通过通信共享内存而不是通过共享内存实现通信。
	如果说 goroutine 是 Go程序并发的执行体，channel 就使他们之间的连接。
	channel 是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
	Go语言中的通道（channel）是一种特殊的类型，通道像一个传送带或者队列，总是遵循先进先出
	的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候
	需要为其指定元素类型。
声明channel：
	格式如下：
		var 变量 chan 元素类型
	channel类型是引用类型
*/

func TestChannel01() {
	//声明通道
	var ch1 chan int
	var ch2 chan string

	//make函数初始化：slice、map、channel
	ch3 := make(chan int, 6)
	fmt.Println(ch1, ch2, ch3)

	//通道的操作：发送、接收、关闭
	//发送接收 <-
	ch3 <- 10 //把 10 发送到ch3
	//<-ch3        //从ch3中接收值，直接丢弃
	ret := <-ch3 //从ch3中接收值，保存到变量ret中
	fmt.Println(ret)
	ch3 <- 9
	ch3 <- 7
	//关闭通道
	close(ch3)
	//1.关闭的通道再接收，如果通道里还有值就能拿到，如果没有值了就取到对应类型的零值
	ret2 := <- ch3
	fmt.Println(ret2)
	//2.不能往关闭的通道发送值，会引发panic
	//ch3 <- 20
	//3.关闭一个已经关闭的通道会引发panic
	//close(ch3)
}
