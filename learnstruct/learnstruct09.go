package learnstruct

import "fmt"

// ------------- 结构体体的匿名字段 -------------------
/*
	结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名称的字段成为匿名字段
	匿名字段类型不能相同
 */

type dog struct{
	string
	int
}


func PrintStruct09() {
	p1 := dog{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.string, p1.int)
}
