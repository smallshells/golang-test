package learnstruct

import "fmt"

// ------------- 嵌套结构体 -------------------
/*
	匿名结构体在嵌套结构体中的应用
 */

type address struct {
	provice string
	city    string
}

//结构体嵌套
type school struct {
	stuNumber int64
	teaNumber int64
	addr      address
}

//结构体匿名嵌套
type company struct {
	name string
	address              //这样要防止冲突
}

func PrintStruct10() {

	sch := school{
		stuNumber: 1000,
		teaNumber: 100,
		addr: address{
			provice: "陕西",
			city:    "西安",
		},
	}
	fmt.Println(sch)
	fmt.Println(sch.stuNumber, sch.addr.city)

	//匿名嵌套访问字段，直接到底层
	com := company{
		name: "alibaba",
		address: address{
			provice: "江苏",
			city:    "杭州",
		},
	}
	fmt.Println(com)
	fmt.Println(com.name, com.city)
}
