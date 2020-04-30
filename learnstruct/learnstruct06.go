package learnstruct

import "fmt"

// ------------- 结构体构造函数 -------------------
/*
	就是一个函数：返回这个结构体的实例
 */

/****learnstruct02.go中定义******
type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}
*/

//构造函数：结构体较小时可以返回值类型
func newPerson01(name string, age int8, gender string, hobby []string) person {
	return person{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}

//构造函数：结构体较大时最好返回指针，节省内存开销
func newPerson02(name string, age int8, gender string, hobby []string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}

func PrintStruct06() {
	ps1 := newPerson01("lin",33,"男",[]string{"编程","看书"})
	fmt.Println(ps1)
	psp1 := newPerson02("qiao",32,"女",[]string{"设计","手机"})
	fmt.Println(psp1)
}
