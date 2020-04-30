package learntype

import "fmt"

//-----------------修改字符串----------------
/*
	Go语言的字符串是不能修改的。
	要修改字符串，需要先将其转换成 []rune 或 []byte ，完成后再转换为 string 。
	无论哪种转换，都会 重新分配内存，并复制字节数组。
*/

func PrintType09() {
	s1 := "big"
	//强制类型转换
	byteS1 := []byte(s1)        //把字符串s1强制转换成一个byte切片
	byteS1[0] = 'p'             //只是修改了切片byteS1中的第一个切片值
	fmt.Println(s1)             //原来的字符串并没有改变
	fmt.Println(string(byteS1)) //把byte切片强制转换成字符串输出

	s2 := "白萝卜"
	runeS2 := []rune(s2)        //把字符串s2强制转换成一个rune切片
	runeS2[0] = '红'             //只是修改了切片byteS2中的第一个切片值
	fmt.Println(s2)             //原来的字符串并没有改变
	fmt.Println(string(runeS2)) //把rune切片强制转换成字符串输出
}
