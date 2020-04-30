package learnstruct

import (
	"encoding/json"
	"fmt"
)

// ------------- 结构体与json -------------------
/*
	1.序列化  ： 把Go语言中的结构体变量转变成json格式的字符串
	2.反序列化： 把json格式的字符串转变成Go语言中的结构体变量
*/

type cow struct {
	Name string `json:"name"`  //用json时Name变成name
	Age  int    `json:"age"`
}

func PrintStruct12() {
	cow1 := cow{
		Name: "牛牛",
		Age:  5,
	}
	//序列化
	p, err := json.Marshal(cow1)
	if err != nil {
		fmt.Println("marshal failed, err: ", err)
	}
	fmt.Printf("%#v\n", string(p))

	//反序列化
	str := `{"name":"小小","age":18}`
	var cow2 cow
	json.Unmarshal([]byte(str),&cow2)
	fmt.Printf("%#v\n",cow2)
}
