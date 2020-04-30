package learnchannel

import "fmt"

//---------------- select 多路复用------------------------

func TestChannel06() {
	var ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case ret := <-ch:
			fmt.Println(ret)
		}
	}
}
