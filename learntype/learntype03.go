package learntype

import "fmt"

//-------复数型-----------

/*
complex64 和 complex128 两种类型
复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位
 */

func PrintType03(){
	c1 := 1 + 2i
	//经过测试 c1默认为complex128
	fmt.Printf("the type of 'c1' is %T\n", c1)
	var c2 complex64 = 1 - 2i  //显式定义complex64
	fmt.Println(c2)
	var c3 complex128 = 1 - 2i
	var c4 complex128
	c4 = c1 + c3  // 虚数计算
	fmt.Println(c4)
}
