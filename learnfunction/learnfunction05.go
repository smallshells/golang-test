package learnfunction

import "fmt"

// ---------------- 匿名函数 ---------------
/*
	匿名函数一般不用在全局，一般用在函数内部声明函数，因为函数内部不能用 func 声明函数
*/

//匿名函数
var f7 = func(x, y int) {
	fmt.Println("外部匿名函数：", x+y)
}

func PrintFunction05() {
	f7(23, 27)

	//函数内部声明匿名函数
	f8 := func(x, y int) {
		fmt.Println("内部匿名函数：", x+y)
	}
	f8(10, 89)

	//如果只是调用一次的匿名函数，还可以简写成立即执行函数
	func(x,y int) {
		fmt.Printf("立即执行函数:  %d\n",x+y)
	}(23,45)
}
