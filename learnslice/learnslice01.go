package learnslice

import "fmt"

// --------------- 切片 ---------------
/*
	切片（slice）是一个拥有相同类型元素的可变长度的序列。
	它是基于数组类型做的一层封装。
	它非常灵活，支持自动扩容。
	切片是一个引用类型，它的内部结构包含 地址、 长度 和 容量 。
	切片一般用于快速地操作一块数据集合。
	声明切片类型的基本语法如下：
		var name []T
	其中：
		name：表示变量名
		T：表示切片中的元素类型

	切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求容量
	切片的容量指底层数组从切面第一个元素到最后一个元素的长度
*/

func PrintSlice01() {
	//切片的定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //true nil相当于没有开创内存空间

	//切片初始化
	s1 = []int{1, 2, 3, 4, 5}
	s2 = []string{"拉萨", "汉中", "西安"}
	fmt.Println(s1, s2)

	//切片长度和容量
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}
