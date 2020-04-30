package learnslice

import "fmt"

// ----------------  make()及切片的本质 -----------------
/*
	如果需要动态的创建一个切片，我们就需要使用内置的 make() 函数，格式如下：
		make([]T, size, cap)
		T:切片的元素类型
		size: 切片中元素的数量
		cap: 切片的容量

	切片的本质就是对底层数组的封装，它包含了三个信息：
		底层数组的指针；
		切片的长度；
		切片的容量。
	例如：a := [8]int{0,1,2,3,4,5,6,7} , 切片 s := a[:5]
		切片指针指向a[0]位置，切片长度是5，容量是8

	切片就是一个框，框住了一块连续的内存。
	切片属于引用类型，真正的数据都是保存在底层数组里的。
*/

func PrintSlice04() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v\tlen(s1) = %d\tcap(s1) = %d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 5)
	fmt.Printf("s2 = %v\tlen(s2) = %d\tcap(s2) = %d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 0, 10)
	fmt.Printf("s3 = %v\tlen(s3) = %d\tcap(s3) = %d\n", s3, len(s3), cap(s3))
}
