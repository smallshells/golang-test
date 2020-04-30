package learnstruct

import "fmt"
// ------------- 自定义类型和类型别名 -------------------
/*
	在Go语言中有一些基本的数据类型，如 字符串型 、整型、浮点型、布尔型 等数据类型，
	Go语言可以使用 type 关键字来定义自定义类型。
	自定义类型是定义了一个全新的类型。
	我们可以基于内置的基本类型定义，也可以通过struct定义。例如：
		type myInt int      将myInt定义为int类型
	通过 type 关键字的定义，myInt 就是一种新的类型，它具有 int 的特性。

	类型别名就是给类型取个别的名字，二者等同。
*/

type myInt int     //自定义类型
type yourInt = int //类型别名  rune = int32 、 byte = uint8 就是类型别名

func PrintStruct01() {
	var n myInt = 9
	var m yourInt = 1
	fmt.Println(n)
	fmt.Println(m)
	fmt.Printf("the type of 'n' is %T\n", n)
	fmt.Printf("the type of 'm' is %T\n", m)
}
