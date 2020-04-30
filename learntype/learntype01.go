package learntype

import "fmt"

//-----------整数及进制-------------

/*
整型分为以下两大类：
按长度分为：int8,int16,int32,int64
对应的无符号整型：uint8,uint16,uint16,uint64
其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型

uint8    无符号8位整型（ 0 到 2e8-1）
uint16   无符号16位整型（ 0 到 2e16-1）
uint32   无符号32位整型（ 0 到 2e32-1）
uint64   无符号64位整型（ 0 到 2e64-1）
int8     有符号8位整型  [ -(2e8+1)/2 到 (2e8-1)/2]
int16    有符号16位整型  [ -(2e16+1)/2 到 (2e16-1)/2]
int32    有符号32位整型  [ -(2e32+1)/2 到 (2e32-1)/2]
int64    有符号64位整型  [ -(2e64+1)/2 到 (2e64-1)/2]
*/

/*特殊整型
uint       32位操作系统上就是uint32,64位操作系统上就是uint64
int        32位操作系统上就是int32,64位操作系统上就是int64
uintptr    无符号整型，用于存放一个指针
*/

//Go语言不能定义二进制数，可以定义十进制、八进制、十六进制，可以按照二进制数输出
func PrintType01() {
	//十进制
	var a int = 10
	fmt.Printf("%d \n", a) //10
	fmt.Printf("%b \n", a) //1010  占位符%b表示二进制

	//八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //77   占位符%o表示八进制
	fmt.Printf("%d \n", b) //63   把八进制数按十进制输出

	//十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) //ff   占位符%x表示十六进制
	fmt.Printf("%X \n", c) //FF   占位符%X表示大写十六进制
	fmt.Printf("%d \n", c) //255  把十六进制数按十进制输出
	fmt.Printf("%o \n", c) //377  把十六进制数按八进制输出

	//查看变量类型
	fmt.Printf("the type of ‘c’ is %T \n", c) //占位符%T用来查看变量类型

	//声明int8类型的变量
	d := int8(9)       //明确指定int8类型，否则就是默认为int类型
	var e int16 = 3463 //完整形式明确指定
	fmt.Printf("the type of 'd' is %T, the value is %d \n", d, d)
	fmt.Printf("the type of 'e' is %T, the value is %d \n", e, e)
}
