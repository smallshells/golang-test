package learnfunction

import "fmt"

//------------- fmt包常用函数 -----------------
/*
	fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。
	向外输出：
		Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化
		输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	获取输入：
		有Scan、Scanf、Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。
*/

func PrintFunction08() {
	//Print例子
	fmt.Print("在终端打印该信息。")
	name := "俊林"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端单独显示一行")

	//Scan
	//var s string
	//fmt.Scan(&s)
	//fmt.Println(s)

	//Scanf
	var (
		name1  string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s", &name1, &age, &class)
	//fmt.Println(name1, age, class)

	//Scanln
	fmt.Scanln(&name1, &age, &class)
	fmt.Println(name1, age, class)
}
