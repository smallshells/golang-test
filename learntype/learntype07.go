package learntype

import (
	"fmt"
	"strings"
)

//--------------字符串常用操作------------------
/*
	len(str)                                  求长度
	+ 或 fmt.Sprintf                          拼接字符串
	strings.Split                             分割
	strings.contains                          判断是否包含
	strings.HasPrefix,strings.HasSuffix       前缀/后缀判断
	string.Index(),strings.LastIndex()        子串出现的位置
	strings.Join(a[]string,sep string)        join操作
 */

func PrintType07(){
	name := "俊林"
	aim  := "要打赢"

	//字符串长度
	fmt.Println(len(name))

	//字符串拼接
	life1 := name + aim
	life2 := fmt.Sprintf("%s%s", name, aim)
	fmt.Println(life1,life2)

	//分割
	path := "\"C:\\Users\\Administrator\\go\""
	ret := strings.Split(path,"\\")
	fmt.Println(ret)

	//包含 返回bool
	fmt.Println(strings.Contains(life1,"俊林"))

	//前缀和后缀判断
	fmt.Println(strings.HasPrefix(life2,"俊"))
	fmt.Println(strings.HasSuffix(life2,"打"))

	//查找子串出现的位置
	s := "abcddhg"
	fmt.Println(strings.Index(s, "d"))

	//最后一次出现的位置
	fmt.Println(strings.LastIndex(s, "d"))

	//join操作
	fmt.Println(strings.Join(ret,"."))
}