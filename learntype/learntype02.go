package learntype

import "fmt"

//-------浮点型-----------

/*
Go语言支持两种浮点型数：float32和float64.这两种浮点型整数格式遵循 IEEE 754 标准：
float32的浮点数的最大范围约为 3.4e38 ,  可以使用常量定义：math.MaxFloat32.
float64的浮点数的最大范围约为 1.8e308 , 可以使用常量定义：math.MaxFloat64.
//Go语言中类型不同，不能相互赋值
*/

func PrintType02() {
	f1 := 1.2345
	//默认Go语言中的小数都是float64类型
	fmt.Printf("the type of 'f1' is %T, the value is %f \n", f1, f1)
	var f2 float32 = 12.345  //显式声明float32类型
	f3:=float32(23.345)      //显式声明float32类型
	fmt.Printf("the type of 'f2' is %T, the value is %f \n", f2, f2)
	fmt.Printf("the type of 'f3' is %T, the value is %f \n", f3, f3)
}
