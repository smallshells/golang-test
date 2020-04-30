package learnstruct

import "fmt"

// ------------- 结构体指针 -------------------
/*
	结构体是值类型
	注意：在结构体中支持对结构体指针直接使用 . 来访问结构体成员。
*/

/****learnstruct02.go中定义******
type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}
*/

//值类型传递，参数是拷贝
func fs(x person) {
	x.name = "hao"
}

//指针类型传递，传的是地址
func fsp(x *person) {
	//(*x).name = "hao" 可以写成下面这样：Go语言的语法糖
	x.name = "hao"  //自动根据指针找对应的变量
}

func PrintStruct03() {
	var ps person
	ps.name = "jun"
	ps.age = 33
	ps.gender = "男"
	ps.hobby = []string{"编程", "读书"}

	fs(ps) //值赋值
	fmt.Println(ps.name)

	fsp(&ps) //指针赋值
	fmt.Println(ps.name)
}
