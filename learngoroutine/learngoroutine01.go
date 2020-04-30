package learngoroutine

import (
	"fmt"
	"sync"
)

//---------------------------------------- 并发 --------------------------------------------

/************************************基础概念*************************************
并发：同一时段同时在做多个事情
并行：同一时刻同时在做多个事情
进程：一个程序启动之后就创建了一个进程
线程：操作系统调度的最小单位
协程：用户态的线程
Go语言的并发通过 goroutine 实现。goroutine类似线程，属于用户态的线程，我们可以根据要创建成千上万
个 goroutine 并发工作，goroutine 是由Go语言的运行时调度完成的，而线程是由操作系统调度完成的。
Go语言还提供 channel 在多个 goroutine 之间进行通信。goroutine 和 channel 是 Go语言秉承的 CSP并
发模式的重要实现基础。
*/
/*************************************goroutine*************************************
在java/C++中我们要实现并发编程的时候，我们通常需要自己维护一个线程池，并且需要自己去包装一个又一个
的任务然后自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量心智。能不能有一种机制，
程序员只需要定义多个任务，让系统去帮助我们把这些任务分配到CPU上实现并发执行呢？Go语言中的goroutine
就是这样一种机制，goroutine的概念类似于线程，但goroutine由Go程序运行时调度和管理。Go程序会智能地将
goroutine中的任务合理的分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经
内置了调度和上下文切换的机制。
Go语言编程不需要自己写进程、线程、协程，你的技能包里只有一个技能goroutine，当你需要让某个任务并发执
行的时候，你只需要起一个goroutine就可以了，就是这么简单粗暴。
Go程序中使用 go 关键字为一个函数创建一个 goroutine。一个函数可以被创建多个 goroutine，
一个goroutine必定对应一个函数。
*/

//利用sync.WaitGroup优雅的等待goroutine结束
var wg sync.WaitGroup   //是一个结构体，它里面有一个计数器

func hello() {
	defer wg.Done() //结束了告诉登记，计数器-1。不管下面代码会不会报错，都要执行
	fmt.Println("Hello 拉萨")
}

func TestGoroutine01() {
	wg.Add(1)  //计数器+1，登记1个goroutine
	go hello() //1.创建一个goroutine 2.在新的goroutine中执行hello函数
	fmt.Println("Hello 汉中")
	//time.Sleep(time.Second)  //生硬的让主程序等一秒
	//等hello执行完（执行hello函数的那个goroutine执行完）
	wg.Wait() //阻塞，一直等待所有的goroutine结束，计数器==0
	fmt.Println("hello jun")
}
