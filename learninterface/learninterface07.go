package learninterface

import (
	"fmt"
)

// ------------------------------ 类型断言 -------------------------------
/*
	空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
	一个接口的值是由 一个具体类型 和 具体类型的值 两部分组成的。这两部分分别称为接口的 动态类型 和 动态值
	想要使用空接口中的值，这个时候就可以使用类型断言，其语法格式：
		x.(T)
	其中：
*/

//断言一：
func assign1(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了！")
	} else {
		fmt.Println(str)
	}
}

//断言二：
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串：", t)
	case int:
		fmt.Println("是一个整型：", t)
	case bool:
		fmt.Println("是一个布尔：", t)
	}
}

func PrintInterface07() {

	assign1(78)

	assign2(34)
	assign2("jun")

}
