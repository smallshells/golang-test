package learnstruct

import "fmt"

// ------------- 结构体定义 -------------------
/*
	Go语言中的基础数据类型可以表示一些事物某个基本属性，但是当我们想表达一个事物的全部或者部分属性时，
	这时候再用单一的基本数据类型明显就无法满足需求了。
	Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这些数据类型叫结构体，英文名称 struct。
	我们可以通过 struct 来定义自己的类型。
	Go语言中通过 struct 来实现面向对象。
	结构体定义：
	使用 type 和 struct 关键字来定义结构体，具体格式如下：
		type 类型名 struct{
			字段名 字段类型
			字段名 字段类型
			...
		}
	其中：
		类型名：标识自定义结构体的名称，在同一个包内不能重复
		字段名：标识结构体字段名，结构体中的字段名必须唯一。
		字段类型：标识结构体字段的具体类型。
*/

type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}

func PrintStruct02() {
	var ps person
	ps.name = "jun"
	ps.age = 33
	ps.gender = "男"
	ps.hobby = []string{"编程","读书"}
	fmt.Printf("The type of 'p' is %T\n", ps)
	fmt.Println(ps)

	//定义匿名结构体:作用于一些临时场景
	var point struct{
		x int
		y int
	}
	point.x = 89
	point.y = 23
	fmt.Printf("The type of 'point' is %T\n", point)
	fmt.Println(point)
}
