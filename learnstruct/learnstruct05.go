package learnstruct

import "fmt"

// ------------- 创建结构体指针 -------------------
/*
	方法一：可以通过使用 new 关键字对结构体进行实例化，得到的是结构体的地址。
	格式如下：
		var p = new(person)
	其中：
		p      : 指针类型
		person : 结构体类型
	方法二：通过初始化得到结构体指针。
	格式如下：
		var p = &person{
			字段：值
			字段：值
			...
		}
	或者：
	var p = &person{
			值，值，值...
		}

	注意：结构体在在内存中是连续的
*/

/****learnstruct02.go中定义******
type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}
*/

func PrintStruct05() {

	var psp1 = new(person) //ps就是一个person结构体的示例，只不过这个示例是指针
	fmt.Printf("type psp1 = %T\npsp1 = %#v\npsp1 = %p\npsp1地址 = %p\n", psp1, psp1, psp1, &psp1)

	//用结构体指针为结构体成员赋值
	psp1.name = "lin"
	psp1.age = 33
	psp1.gender = "男"
	psp1.hobby = []string{"编程", "读书"}
	fmt.Println(*psp1)

	//通过初始化得到结构体指针
	//结构体初始化1：key-value
	var psp2 = &person{
		name:   "qiao",
		age:    32,
		gender: "女",
		hobby:  []string{"平面设计", "玩手机"},
	}
	fmt.Println(*psp2)

	//结构体初始化2：使用值列表，值列表和声明的顺序要一致
	psp3 := &person{
		"fan", 2, "男", []string{"外出", "喝奶"},
	}
	fmt.Println(*psp3)
}
