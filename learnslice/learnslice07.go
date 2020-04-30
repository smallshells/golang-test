package learnslice

import "fmt"

// ----------------- 切片追加元素 -----------------
/*
	利用 append() 函数
	Go语言有自己的扩容策略，可以查看 slice.go 源码
*/

func PrintSlice07() {
	s1 := []string{"拉萨", "成都", "汉中"}
	fmt.Printf("s1 = %v\tlen(s1) = %d\tcap(s1) = %d\n", s1, len(s1), cap(s1))

	//append() 函数必须用原来的切片变量接收返回值
	s1 = append(s1, "西安")
	fmt.Printf("s1 = %v\tlen(s1) = %d\tcap(s1) = %d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "北京", "广州")
	fmt.Printf("s1 = %v\tlen(s1) = %d\tcap(s1) = %d\n", s1, len(s1), cap(s1))

	ss := []string{"乌鲁木齐", "榆林","渭南"}

	s1 = append(s1, ss...)   //ss... 就是把ss拆开成字符串   ... 表示拆开
	fmt.Printf("s1 = %v\tlen(s1) = %d\tcap(s1) = %d\n", s1, len(s1), cap(s1))

	var s2 []int
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		fmt.Printf("len(s2):%d\tcap(s2):%d\tptr:%p s2:%v\n", len(s2), cap(s2), s2, s2)
	}
}
