package learnlock

import (
	"fmt"
	"sync"
)

// ------------ 互斥锁 -------------------
/*

 */

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

//没有互斥锁的结果 x可能不为10000
func TestLock011() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

//有互斥锁的版本
func TestLock012() {
	wg.Add(2)
	go addLock()
	go addLock()
	wg.Wait()
	fmt.Println(x)
}

func addLock() {
	lock.Lock() // 加锁
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	lock.Unlock() // 解锁
	wg.Done()
}

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
