package learnstruct

import "fmt"

// ------------- 方法和接受者 -------------------
/*
	Go语言中的 方法Method 是一种作用于特定类型变量的函数。
	这种特定类型变量叫 接收者Receiver 。
	接收者的概念就类似于其他语言中的 this 或者 self

	方法格式如下：
		func (接收者变量 接收者类型) 方法名(参数列表)(返回参数){
			函数体
		}
	其中：
		接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，
				   而不是 self 、 this 之类的命名。例如， Person类型的接收者变量应该命名为 p ，
				   Connector类型的接收者变量应该命名为 c 等。
		接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
		方法名、参数列表、返回参数：具体格式与函数定义相同。

	注意：指针接受者使用时机：
		1.需要修改接收者中的值。
		2.接收者是拷贝代价比较大的对象。
		3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
 */

/*结构体：learnstruct02.go中定义
type person struct {
	name   string
	age    int8
	gender string
	hobby  []string
}
*/

/*构造函数：在learnstruct06.go中定义
func newPerson02(name string, age int8, gender string, hobby []string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}
 */

//(p person)是接收者，这是值接收者：传拷贝
func (p person) run(){
	fmt.Printf("%s:正在跑步\n",p.name)
}

//(p *person)是接收者，这是指针接收者:传地址
func (p *person) guonian(){
	p.age++
}

func PrintStruct07() {
	var psp = newPerson02("lin",33,"男",[]string{"编程","读书"})
	psp.run()
	psp.guonian()
	fmt.Println(psp.age)
}
