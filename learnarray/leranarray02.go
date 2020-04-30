package learnarray

import "fmt"

// --------------- 数组初始化 ---------------
/*
	方法一：初始化数组可以使用初始化列表来设置数组元素的值
	方法二：根据初始值自动推断数组的长度是多少
	方法三：根据索引来初始化
*/

func PrintArray02() {
	//方法一
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化,第三个默认0
	var cityArray = [3]string{"拉萨", "汉中", "西安"} //使用指定的初始值完成初始化
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	//方法二
	xArray := [...]int{1,2,45,65,23,67,8}
	fmt.Println(xArray)

	//方法三
	indexArray := [5]int{0:1,4:8}
	fmt.Println(indexArray)
}
