package learnvar

import "fmt"

//--------------变量定义--------------

/*
变量声明的三种方式
var name string = "hao"   第一种（完整写法）
var name = "hao"          第二种（类型推导）
name := "hao"             第三种（短变量声明）只能在函数内部使用
*/
//函数外的语句都必须以关键字开始（var,const,func等）
var (
	name string
	age  int
	isOk bool
)

func PrintVar() {
	//var sex string = "男"   局部变量声明了就要使用，不然编译不能通过
	name = "俊林"
	age = 33
	isOk = false
	fmt.Println(isOk)
	fmt.Printf("name:%s\ngae:", name)
	fmt.Println(age)
}
