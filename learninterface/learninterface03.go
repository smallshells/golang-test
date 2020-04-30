package learninterface

import "fmt"

//-------------------- 接口的实现 -----------------------
/*
	接口变量保存分为两部分：
		1.值的类型
		2.值
	这样就实现了任一个接口变量能够存储不同的值
*/
type animal interface {
	move()
	eat(string)
}

type cat02 struct {
	name string
	feet int8
}

func (c cat02) move() {
	fmt.Println("走猫步")
}

func (c cat02) eat(food string) {
	fmt.Println("小猫吃", food)
}

type chicken02 struct {
	feet int8
}

func (c chicken02) move() {
	fmt.Println("鸡动")
}

func (c chicken02) eat(food string) {
	fmt.Println("鸡吃", food)
}

func PrintInterface03() {
	var a1 animal //定义一个接口类型的变量
	fmt.Printf("the type of a1 is %T\n",a1)
	bc := cat02{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	fmt.Println(a1)
	a1.eat("小黄鱼")
	fmt.Printf("the type of a1 is %T\n",a1)
	fmt.Printf("the value of a1 is %#v\n",a1)

	kfc := chicken02{
		feet:2,
	}

	a1 = kfc
	a1.eat("大米")
	fmt.Printf("the type of a1 is %T\n",a1)
}
