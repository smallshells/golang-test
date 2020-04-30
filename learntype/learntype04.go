package learntype

import "fmt"

//----------------布尔值--------------
/*
Go语言中以 bool 类型进行声明布尔型数据，布尔型数据只有 true（真）和 false（假）两个值
注意：
	1.布尔类型变量的默认值为false
	2.Go语言中不允许将整型强制转换为布尔型
	3.布尔型无法参与数值计算，也无法与其他类型进行转换
*/

func PrintType04() {
	b1 := true
	var b2 bool  //默认值为false
	fmt.Printf("the type of 'b1' is %T, value : %v \n", b1, b1)
	//占位符%v表示打印变量的值，不管什么类型都可以打印
	fmt.Printf("the type of 'b2' is %T, value : %v \n", b2, b2)
}
