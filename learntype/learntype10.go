package learntype

import (
	"fmt"
	"math"
)

//---------------类型转换------------------
/*
	Go语言中只有强制类型转转，没有隐式类型转换。
	该语法只能在两个类型之间支持相互转换的时候使用。
	强制类型转换的基本语法如下：
		T(表达式)
	其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等。
*/

func PrintType10() {
	var a, b = 3, 4
	var c int
	//math.Sqrt()接受的参数是float类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))   //结果强制转换成int类型
	fmt.Println(c)
}
