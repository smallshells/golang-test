package learninterface

import "fmt"

//---------------- 接口和类型的区别 ------------------
/*
	同一个结构体可以实现多个接口
	同一个接口可以被多个结构体实现
	接口可以嵌套
*/

type Animal interface {
	mover   //结构体嵌套，相当于move()
	eater   //结构体嵌套，相当于eat()
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type dog04 struct {
	name string
}

func (d dog04) move() {
	fmt.Println(d.name, "可以动")
}

func (d dog04) eat() {
	fmt.Println(d.name, "爱吃骨头")
}

type cat04 struct {
	name string
}

func (c cat04) move() {
	fmt.Println(c.name, "可以动")
}

func (c cat04) eat() {
	fmt.Println(c.name, "爱吃鱼")
}

func PrintInterface05() {

}
