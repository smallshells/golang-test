package learnfunction

import "fmt"

// ---------------- 内置函数简介 ---------------
/*
	close  		主要用来关闭channel
	len			用来求长度，比如string, array, slice, map, channel
	new			用来分配内存，主要用来分配值类型，比如int,struct.返回的是指针
	make		用来分配内存，主要用来分配引用类型，比如chan, map, slice.
	append		用来追加元素到数组、slice中
	panic	 	抛出一个错误，程序退出
	recover     尝试恢复错误
*/

//情况一：正常
func funcA1() {
	fmt.Println("A")
}

func funcB1() {
	fmt.Println("B")
}

func funcC1() {
	fmt.Println("C")
}

func mode01() {
	funcA1()
	funcB1()
	funcC1()
}

//情况二：出现错误，程序崩溃退出
func funcA2() {
	fmt.Println("A")
}

func funcB2() {
	panic("程序出现了严重错误") //程序崩溃退出
	fmt.Println("B")
}

func funcC2() {
	fmt.Println("C")
}

func mode02() {
	funcA2()
	funcB2()
	funcC2()
}

//情况三：出现错误后程序崩溃前处理一些情况：比如释放数据库等
func funcA3() {
	fmt.Println("A")
}

func funcB3() {
	defer func() {
		fmt.Println("释放数据库连接")
	}()
	panic("程序出现了严重错误") //程序崩溃退出
	fmt.Println("B")
}

func funcC3() {
	fmt.Println("C")
}

func mode03() {
	funcA3()
	funcB3()
	funcC3()
}

//情况四：尝试恢复错误
func funcA4() {
	fmt.Println("A")
}

func funcB4() {
	defer func() {	//defer一定要在可能引发panic的语句之前定义
		err := recover()   //尽量少用，该panic还是要panic. recover()必须搭配defer使用
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()

	panic("程序出现了严重错误") //程序崩溃退出
	fmt.Println("B")
}

func funcC4() {
	fmt.Println("C")
}

func mode04() {
	funcA4()
	funcB4()   //函数内部出现了错误，但尝试了恢复，不影响后面的程序执行
	funcC4()
}

//主调用程序
func PrintFunction07() {
	mode04()
}
