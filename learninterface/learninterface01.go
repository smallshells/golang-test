package learninterface

import "fmt"

//------- 初识接口 ----------
/*
	接口就是一种规范，就像车的规范一样。只要能开走就能算做车，比如：汽车、自行车、摩托车等，都是车。
 */

type speaker interface {
	speak()
}

type cat01 struct{}

type dog01 struct{}

type person01 struct{}

func (c cat01) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog01) speak() {
	fmt.Println("汪汪汪~")
}

func (p person01) speak() {
	fmt.Println("小样~")
}

func da(x speaker) {
	x.speak()
}

func PrintInterface01() {
	var c1 cat01
	var d1 dog01
	var p1 person01

	da(c1)
	da(d1)
	da(p1)
}
