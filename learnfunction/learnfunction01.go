package learnfunction

import "fmt"

//------------------ 函数定义 -----------------
/*
	Go语言中定义函数使用 func 关键字，具体格式如下：
		func 函数名(参数)(返回值){
			函数体
		}
	其中：
		函数名：由字母、数字、下划线。但函数的第一个字母不能是数字。在同一个包内，函数名不能重名。
		参数：参数由参数变量和参数变量的类型组成，多个参数之间使用 , 分割。没有默认参数
		返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值类型，
			   多个返回值必须用 () 包裹，并用 , 分隔。
			   如果使用了返回值变量，相当于提前在函数内部声明了一个返回值变量，内部可以直接调用
		函数体：实现指定功能的代码块。
*/

//定义函数
func sum(x int, y int) (ret int) {
	return x + y
}

//-----下面代码等同于上面------
func sum1(x int, y int) (ret int) {      //使用命名的返回值相等于在函数内部先声明了一个返回值变量
	ret =  x + y
	return
}

//没有返回值有参数的函数定义
func printSum(x int, y int) {
	fmt.Println(x + y)
}

//没有参数有返回值的函数定义
func getName() (name string) {
	return "getName"
}

//没有参数没有返回值的函数定义
func printName() {
	fmt.Println("printName")
}

//多个返回值
func getNu() (int, int) {
	return 3, 4
}

//参数的类型简写:类型一致时
func getSum(x, y int) int {
	return x + y
}

//可变长参数
func turnPar(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)   // y 的类型是一个切片
}

func PrintFunction01() {
	r := sum(2, 3)
	fmt.Println(r)

	printSum(2, 3)

	fmt.Println(getName())

	printName()

	fmt.Println(getNu())

	fmt.Println(getSum(2, 3))

	turnPar("打雷了",12,3,4,5)
	turnPar("junlin")
}
