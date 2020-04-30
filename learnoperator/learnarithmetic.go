package learnoperator

import "fmt"

// ----------------- 算术运算符 -----------------
/*
	+  相加
	-  相减
	*  相乘
	/  相除
	%  求余

	注意： ++（自增）和 --（自减）在Go语言中是单独的语句，并不是运算符
*/

func PrintArithmetic() {
	var (
		a = 5
		b = 2
	)
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	a++                      //这是一个单独的语句，不能当成值写成 a = a++ 这样的形式
	fmt.Println("a++ : ", a) //不能写成fmt.Println(a++)
}
