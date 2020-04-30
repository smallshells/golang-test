package learnarray

import "fmt"

// --------------- 数组遍历 ---------------
/*
	数组可以通过下标进行访问，下标是从 0 开始，最后一个元素下标是 len-1 ，访问越界，则处罚访问越界，会panic
	有两种方法
*/

func PrintArray03() {
	var cityArray = [...]string{"拉萨", "汉中", "西安"}

	//方法一：包含汉字的字符串用下标遍历就要出错
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	//方法二:
	for index, value := range cityArray {
		fmt.Println(index, value)
	}
}
