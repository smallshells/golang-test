package learnfunction

import (
	"fmt"
	"strings"
)

// ---------------- 闭包 ---------------
/*
	闭包指的是一个函数和与其相关的引用环境组合而成的实体。
	简单来说，闭包 = 函数 + 引用环境。
	底层原理：
		1.函数可以作为返回值；
		2.函数内部查找变量的顺序，先局部后全局。
	闭包目的：
		简单说就是创造一个函数利用外部引用做点事情：
			示例一：创造一个函数让 外部引用 相加 某个变量
			示例二：创造一个函数让 外部引用 添加到 某个变量（字符串）后面
			示例三：创造一个函数让 外部引用 相加或相减 某个变量
			示例四：创造一个函数让 外部引用 执行
*/

/*-------------------  示例一：  -----------------
变量 f1 是一个函数并且它引用了其外部作用域中的 x 变量，此时 f1 就是一个闭包。
在 f1 的声明周期内，变量 x 一直有效。
*/
func adder(x int) func(int) int {
	return func(y int) int { //匿名函数包含了外部变量 x，相当于是对 x 的重新包装
		x += y
		return x
	}
}

/*-------------------  示例二：  -----------------
创造一个给字符串结尾添加固定字符串的函数
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

/*-------------------  示例三：  -----------------
基础数值加减函数
*/
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

/*-------------------  示例四：  -----------------
将f11用f12包装后，传递给f10
*/
func f10(x func()) {
	fmt.Println("f10")
	x()
}

func f11(x, y int) {
	fmt.Println("f11")
	fmt.Println(x, y)
}

func f12(x func(int, int), a, b int) func() {
	fmt.Println("f12")
	return func() { //匿名函数包含了外部变量 x,a,b 相当于是对它们的重新包装
		x(a, b)
	}
}

func PrintFunction06() {
	//示例一
	var f1 = adder(10)
	fmt.Println(f1(10))
	fmt.Println(f1(20))
	fmt.Println(f1(30))
	f2 := adder(100)
	fmt.Println(f2(40))
	fmt.Println(f2(100))

	//示例二
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test.txt"))

	//示例三
	fadd, fsub := calc(10)
	fmt.Println(fadd(1), fsub(2))
	fmt.Println(fadd(3), fsub(4))
	fmt.Println(fadd(5), fsub(6))

	//示例四
	f3 := f12(f11, 2, 3)
	f10(f3)
}
