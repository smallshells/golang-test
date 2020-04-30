package learninterface

import "fmt"

//---------------- 值接收者和指针接收者实现接口的区别 ------------------
/*
	使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存。
	指针接收者实现接口只能存结构体指针类型的变量
*/

type Mover interface {
	move()
}

type dog03 struct {}

//使用值接收者
//func (d dog03) move(){
//	fmt.Println("running...")
//}

//使用指针接收者
func (d *dog03) move(){
	fmt.Println("running...")
}

func PrintInterface04() {
	var mm Mover

	d1 := dog03{}
	d2 := &dog03{}

	//值接收时都可以，指针接收者时要加取地址符&
	mm = &d1
	fmt.Println(mm)
	mm = d2
	fmt.Println(mm)

}
