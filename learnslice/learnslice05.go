package learnslice

import "fmt"

// ------------ 切片不能直接比较 -----------------
/*
	切片之间是不能比较的，我们不能使用 == 操作符判断两个切片是否含有全部相等元素。
	切片唯一合法的比较操作符是和 nil 比较。
	一个 nil 值得切片并没有底层数组，一个 nil 值得切片的长度和容量都是 0.
	但一个长度和容量都是 0 的切片不一定是 nil
	要判断一个切片是否为空，要用 len(s) == 0 来判断，不应该使用 s == nil 来判断
 */

func PrintSlice05(){
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	fmt.Printf("len(s1) = %d, len(s1) = %d, 's1 == nil' is %v\n", len(s1), cap(s1), s1==nil)
	fmt.Printf("len(s2) = %d, len(s2) = %d, 's2 == nil' is %v\n", len(s2), cap(s2), s2==nil)
	fmt.Printf("len(s3) = %d, len(s3) = %d, 's3 == nil' is %v\n", len(s3), cap(s3), s3==nil)
}

