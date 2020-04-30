package learnchannel

import "fmt"

//------------------------ 无缓冲通道和有缓冲通道 -----------------------------
/*
	无缓冲通道相当于4x100米接力，交棒的人必须有接棒的才行，不然就交不出去，形成阻塞。
	缓冲通道里面有缓冲容量，可以先把值放进缓冲区。可以使程序实现异步操作。
*/

func recv(ch chan bool){
	ret := <- ch  //无缓冲通道会阻塞
	fmt.Println(ret)
}


func TestChannel02() {
	ch1 := make(chan bool)  //无缓冲通道
	go recv(ch1)
	ch1 <- true
	fmt.Println("main函数结束")

	ch2 :=make(chan bool, 1)  // 有缓冲通道
	//len 获取数据量，cap获取容量
	fmt.Println(len(ch2),cap(ch2))
	ch2 <- false
	go recv(ch2)
	go recv(ch2)
	ch2 <- true
}
