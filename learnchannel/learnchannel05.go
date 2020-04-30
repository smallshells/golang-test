package learnchannel

import (
	"fmt"
	"time"
)

//---------------- select 多路复用------------------------

var ch1 = make(chan string, 100)
var ch2 = make(chan string, 100)

func f1(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f1 : %d", i)
		time.Sleep(time.Second)
	}
}

func f2(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("f2 : %d", i)
		time.Sleep(time.Second * 2)
	}
}

func TestChannel05() {
	go f1(ch1)
	go f2(ch2)
	var ret string
	for {
		select {
		case ret = <-ch1:
			time.Sleep(time.Second)
			fmt.Println(ret)
		case ret = <-ch2:
			time.Sleep(time.Second)
			fmt.Println(ret)
		default:
			time.Sleep(time.Second)
			fmt.Println("暂时没有取到值")
		}
	}
}
