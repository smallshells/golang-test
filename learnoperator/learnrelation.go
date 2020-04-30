package learnoperator

import "fmt"

// ------------------ 关系运算符 ------------------
/*
	==  是否相等
	!=  是否不相等
	>   是否大于
	>=  是否大于等于
	<   是否小于
	<=  是否小于等于

	注意：Go语言是强类型，相同类型才能比较
 */

func PrintRelation(){
	var (
		a = 5
		b = 2
	)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)
}