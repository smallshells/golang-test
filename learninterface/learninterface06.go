package learninterface

import "fmt"

//------------ 空接口 -----------------
/*
	空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口
	空接口的变量可以存储任意类型的变量。
	如果想实现一个函数不管什么变量都能传进来，就应该设置成空接口类型
*/

//一般情况下空接口不起名字，直接使用就行了
type x interface{}

//空接口作为函数的参数,任意类型都能接收
func show(a  interface{}) {
	fmt.Printf("The type is %T, the value is %v\n", a, a)
}

func PrintInterface06() {

	//空接口，没有必要起名字，通常定义成下面的格式：
	var xx interface{}
	xx = "jun"
	fmt.Println(xx)
	xx = 98
	fmt.Println(xx)

	//用空接口实现储存多种类型的map
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 10)
	m1["name"] = "俊林"
	m1["age"] = 33
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱歌","写作","编程"}
	fmt.Println(m1)

	//show()函数接受任意类型参数
	show(76)
	show("jun")
	show([...]int{12,12,23,45})
	show([]string{"jun","琳","OK"})
	show(map[int]string{110:"公安",120:"急救",119:"火警"})
	show(struct{x int;y int}{12,13})
}
