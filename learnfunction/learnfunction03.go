package learnfunction

import "fmt"

//---------- 函数作用域 ------------------
/*
	全局变量：定义在函数外部的变量，它在程序整个运行周期内都有效。在函数中可以访问到全局变量
	局部变量：函数内部定义的变量 或者 语句块内部定义的变量，只在内部可以访问。
	函数中查找变量的顺序：先局部后全局
*/

//定义全局变量 num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num = %d\n", num)
}

func testLocalVar01() {
	name := "俊林" //定义局部变量
	fmt.Println(name)
}

func testLocalVar02(x, y int) {
	fmt.Println(x, y) //函数的参数也只在函数内部生效
	if x > 0 {
		z := 100
		fmt.Println(z) //变量 z 只在 if 语句块中有效
	}
	//此处无法使用 fmt.Println(z)
}

func PrintFunction03() {
	testGlobalVar()
	testLocalVar01()
	testLocalVar02(1, 5)
}
