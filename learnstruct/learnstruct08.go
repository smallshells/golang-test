package learnstruct

import "fmt"

// ------------- 任意类型添加方法 -------------------
/*
	不能改别的包中的类型添加方法，只能给自己的包里添加方法
	要想给别的包中的类型添加方法，可以根据那个类型自定义自己的类型，再添加方法
 */
/*自定义类型：在learnstruct01.go中定义
type myInt int     //自定义类型
 */

func (m myInt)hello(){
	fmt.Println("hello，my name is myInt")
}

func PrintStruct08() {
	var mi myInt = 10
	mi.hello()
}
