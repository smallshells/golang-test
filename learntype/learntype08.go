package learntype

import "fmt"

//-----------------遍历字符串------------------------
/*
	byte和rune类型
	组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。
	字符用单引号（'）包裹起来，如：
		var a := '中'
    	var b := 'x'
	Go语言的字符有以下两种：
		1.unit8类型，或者叫byte型，代表了ASCII码的一个字符。
		2.rune类型，代表一个UTF-8字符。Go语言默认为rune。
	当需要处理中文、日文或者其他复合字符时，则需要用到rune类型，rune类型实际是一个int32.
	Go使用了特殊的rune类型来处理Unicode，让基于Unicode的文本处理更为方便，
	也可以使用byte型进行默认字符串处理，性能和扩展性都有照顾。
*/

func PrintType08() {
	s := "hello俊林"
	n := len(s) //求字符串长度，为字符串的byte字节数量
	fmt.Println(n)

	//当字符串有汉字时，len()的值大于字符串中的字符个数，遍历会出错溢出
	//用s[i]这样的方式遍历字符串时，遍历出来的是字符串的byte位数值，对于汉字不能输出
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i]) //%v:值输出；%c：字符输出
	}
	fmt.Println()
	for _, c := range s { //从字符串中拿出具体的字符
		fmt.Printf("%v(%c)\n", c, c) //%v:值输出；%c：字符输出
	}
}
