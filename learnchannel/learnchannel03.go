package learnchannel

import "fmt"

//---------------------- 接收值时判断通道是否关闭 --------------------------
/*
	1.使用 value, ok := <- channel 取值方式，当channel关闭的时候，ok = false
	2.利用for range 循环到通道中取值（推荐）
*/

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)   //不像打开的文件必须在代码中显式的关闭。
}

func TestChannel03() {
	var ch1 = make(chan int, 100)
	go send(ch1)
	//1.利用for循环到通道中取值
	for {
		ret, ok := <-ch1
		if !ok { //判断通道是否关闭
			break
		}
		fmt.Println(ret)
	}

	var ch2 = make(chan int, 50)
	go send(ch2)
	//2.利用for range 到通道中取值
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
