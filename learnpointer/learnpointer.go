package learnpointer

import "fmt"

// ---------------------- 指针 --------------------
/*
	Go语言指针只有两个操作
		1. &：取地址
		2. *：根据地址取值

	对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	指针变量的值是内存地址。
	对指针进行取值（*）操作，可以获得指针变量指向的原变量的值。

	用new()函数开辟一个内存空间：
		var p = new(int)

	make 和 new 的区别：
		1.make 和 new 都是用来申请内存的
		2.new 很少用，一般用来给基本数据类型申请内存，string、 int， 返回的是对应类型的指针
		3.make 是用来给 slice、 map、 chan 申请内存的，make 函数返回的是对应的三个类型本身
 */

func PrintPointer(){
	var p *int  //指针声明，此时 p 是 nil 没有开创内存地址，不能给*p赋值
	n := 89
	fmt.Println(&n)  //打印变量 n 的地址

	p = &n    //将n的地址赋值给指针p
	fmt.Println(*p)   //打印指针p对应的地址的值

	//用new()函数开辟一个内存空间
	var a = new(int)
	*a = 100
	fmt.Println(*a)
}