package learnstruct

import "fmt"

// ------------- 结构体的“继承” -------------------
/*
	Go语言中使用结构体也可以实现其他编程语言中面向对象的继承
 */

//动物
type Animal struct{
	name string
}

func (a *Animal) move(){
	fmt.Println(a.name,"会动")
}

//狗
type Dog struct {
	Feet int8
	*Animal   //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang(){
	fmt.Println(d.name,"会汪汪汪……")
}

func PrintStruct11() {
	d1 := &Dog{
		Feet:4,
		Animal:&Animal{name:"乐乐"},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
