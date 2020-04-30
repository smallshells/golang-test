package learngoroutine

import (
	"fmt"
	"runtime"
	"sync"
)

//---------------------------------------- goroutine与线程 --------------------------------------------
/*
可增长的栈：
	OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB），一个goroutine的栈在其生命周期开始时只有很小的
	栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB
	虽然极少会用到这么大。所以在Go语言中一次创建十万左右的goroutine也是可以的。
goroutine调度：
	OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为
	m:n调度的技术（复用/调度m个goroutine到n个OS线程）。goroutine的调度不需要切换内核语境，所以调用一个goroutine
	比调度一个线程成本低很多。
GOMAXPROCS:
	Go运行时的调度器使用 GOMAXPROCS 参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
	例如 在一个8核心得机器上，调度器会把Go代码同时调度到8个OS线程上，（GOMAXPROCS是m:n调度中的n）。
	Go语言可以通过 runtime.GOMAXPROCS() 函数设置当前程序并发时占用的CPU逻辑核心数。
	GO1.5版本之前，默认使用的是单核心执行，Go1.5版本之后，默认使用全部的CPU逻辑核心数。
	我们可以通过将任务分配到不同的CPU逻辑核心上实现并行的效果。
系统线程和goroutine关系：
	1.一个操作系统线程对应用户态多个goroutine
	2.Go程序可以同时使用多个操作系统线程。
	3.goroutine和OS线程是多对多得关系，即m:n。
*/

var ww sync.WaitGroup

func a(){
	defer ww.Done()
	for i := 1; i<100; i++ {
		fmt.Println("A:",i)
	}
}

func b(){
	defer ww.Done()
	for i := 1; i<100; i++ {
		fmt.Println("B:",i)
	}
}

func TestGoroutine02() {
	//系统是4核心，Go语言默认开4个
	runtime.GOMAXPROCS(1)  //设置我的Go程序只用一个逻辑核心 m:n中设置n为1。其实就是一个程序跑完了再跑另一个
	ww.Add(2)
	go a()
	go b()
	ww.Wait()
}
