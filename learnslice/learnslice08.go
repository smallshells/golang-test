package learnslice

import "fmt"

//--------------- 复制切片 ------------------
/*
	使用cope()函数复制切片：将一个切片的数据复制到另一个切片空间中。
	使用格式：
		copy(destSlice, srcSlice []T)
		其中：
			srcSlice:数据来源切片
			destSlice:目标切片

*/

func PrintSlice08() {
	s1 := []int{1, 3, 5}
	s2 := s1     //赋值
	var s3 []int //这样声明变量 s3 为 nil
	copy(s3, s1) //由于s3 为nil s1 copy不进去
	fmt.Println(s1, s2, s3)
	var s4 = make([]int, 3, 5)
	copy(s4, s1)
	fmt.Println(s1, s2, s4)

	//改变s1不影响s4，影响s2
	s1[0] = 100
	fmt.Println(s1, s2, s4)

	//改变s4，不影响s1
	s4[1] = 200
	fmt.Println(s1, s2, s4)
}
