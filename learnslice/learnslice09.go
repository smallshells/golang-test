package learnslice

import "fmt"

//--------------- 删除切片中元素 ---------------
/*
	Go语言中没有专门删除切片元素的方法，可以利用 append() 函数来实现:
	要从切片s中删除索引为index的元素，操作方法是：
		s = append(a[:index], a[index+1]...)
*/

func PrintSlice09() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[:]
	//删除索引为2的元素
	s = append(s[:2], s[3:]...)
	fmt.Println(s)
	fmt.Println(a)    //底层数组被修改了
}
