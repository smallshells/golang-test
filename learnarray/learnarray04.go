package learnarray

import "fmt"

// -------------- 多维数组 --------------
/*
	基本操作和一位数组类同
*/

func PrintArray04() {
	//多维数组定义
	var arr [3][2]int

	//多维数组初始化
	arr = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(arr)

	//多维数组遍历
	for i1, v1 := range arr {
		for i2, v2 := range v1 {
			fmt.Printf("arr[%d][%d] = %d\t", i1, i2, v2)
		}
		fmt.Println()
	}

}
