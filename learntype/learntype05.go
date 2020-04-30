package learntype

import "fmt"

//--------------常用占位符-------------

func PrintType05() {
	var n = 100
	fmt.Printf("%T\n", n) //打印类型
	fmt.Printf("%v\n", n) //打印值
	fmt.Printf("%b\n", n) //打印二进制
	fmt.Printf("%d\n", n) //打印十进制
	fmt.Printf("%o\n", n) //打印八进制
	fmt.Printf("%x\n", n) //打印十六进制
	var s = "hello LaSa"
	fmt.Printf("%s\n", s)  //打印字符串
	fmt.Printf("%v\n", s)  //打印值
	fmt.Printf("%#v\n", s) //打印值，#表示打印时带着类型符号：字符串类型为""，八进制前面带0
	var c = 1 + 4i
	fmt.Printf("%v\n", c) //打印值
	var b = true
	fmt.Printf("%t\n", b)  //打印布尔值
	var f = 12.34
	fmt.Printf("%f\n", f)  //打印浮点数
	var p = &f
	fmt.Printf("%p\n", p)  //打印浮点数
}
