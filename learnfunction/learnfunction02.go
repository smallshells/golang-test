package learnfunction

import "fmt"

// ----------------- 函数 defer 语句 ------------------
/*
	Go语言中的 defer 语句会将其后面跟随的语句进行延迟处理。
	在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 定义的逆序进行执行。
	也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
	defer 多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接等）

	defer执行时机：
		Go语言的函数中 return 语句在底层并不是原子操作，它分为给 返回值赋值 和 RET指令 两步。
		defer 语句执行的时机就是在返回值赋值操作之后，RET 指令执行之前。
	defer语句在第一次遇到时会把语句中的变量用变量中的值都替代，等待回头执行： 例如 testDefer07() 能很清楚实验
*/

func testDefer01() {
	fmt.Println("start")
	defer fmt.Println(1) //最先 defer 最后被执行
	defer fmt.Println(2)
	defer fmt.Println(3) //最后 defer 相较于其他 defer 最先执行
	fmt.Println("end")
}

//经典 defer 时机例题
func testDefer02() int {
	x := 5
	defer func() {
		x++ //修改的是 x 不是返回值
	}()
	return x //1.把x的值拷贝给返回值（返回值没有命名），返回值=5；2.执行defer修改x，返回值还是5；3.返回返回值5
}
func testDefer03() (x int) {
	defer func() {
		x++ //修改了返回值 x
	}()
	return 5 //1.返回值x=5；2.执行defer x++，返回值x=6；3.返回返回值6
}
func testDefer04() (y int) {
	x := 5
	defer func() {
		x++ //修改的是 x 不是返回值 y
	}()
	return x //1.把x的值拷贝给返回值y，返回值y=5；2.执行defer修改x，返回值y=5；3.返回返回值5
}
func testDefer05() (x int) {
	defer func(x int) {
		x++ //函数内部的 x，不是返回值x
	}(x)
	return 5 //1.把x的值拷贝给返回值，返回值=5；2.执行defer修改函数内部x，外部返回值x=5；3.返回返回值5
}

func testDefer06() (x int) {
	defer func(x *int) {
		*x++
	}(&x)    //传递x的地址
	return 5 //1.返回值=x=5； 2.defer x=6 ; 3.RET返回6
}

func calcDefer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func testDefer07() {
	a := 1
	b := 2
	defer calcDefer("1", a, calcDefer("10", a, b))
	a = 0
	defer calcDefer("2", a, calcDefer("20", a, b))
	b = 1
}

func PrintFunction02() {
	testDefer01()
	fmt.Println(testDefer02())
	fmt.Println(testDefer03())
	fmt.Println(testDefer04())
	fmt.Println(testDefer05())
	fmt.Println(testDefer06())
	testDefer07()

}
