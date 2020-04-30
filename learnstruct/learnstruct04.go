package learnstruct

import "fmt"

// ------------- 结构体初始化 -------------------
/*
	结构体初始化有两类：
	第一类：先定义再赋值。
	比如：
		var ps person;ps.name = "lin"
	或者：
		psp := new(person);psp.name = "lin"
	第二类：定义时直接赋值：两种格式：
	格式如下：
		var ps = person{
			字段：值
			字段：值
			...
		}
	或者：
	var ps = person{
			值，值，值...
		}
*/

/****learnstruct02.go中定义******
type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}
*/

func PrintStruct04() {

	//结构体出事化1：key-value
	var ps1 = person{
		name:   "qiaoqiao",
		age:    32,
		gender: "女",
		hobby:  []string{"平面设计", "玩手机"},
	}
	fmt.Println(ps1)

	//结构体初始化2：使用值列表，值列表和声明的顺序要一致
	ps2 := person{
		"xiufan", 2, "男", []string{"外出", "喝奶"},
	}
	fmt.Println(ps2)
}
