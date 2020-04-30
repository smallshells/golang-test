package learnslice

import "fmt"

// ------------- 切片遍历 ----------------
/*
	切片的遍历方式和数组是一致的，支持索引遍历和 for range 遍历
*/

func PrintSlice06() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}
