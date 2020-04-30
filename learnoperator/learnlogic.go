package learnoperator

import "fmt"

// ------------------ 逻辑运算符 -----------------
/*
	&&  逻辑 AND 运算符
	||  逻辑 OR 运算符
	！	逻辑 NOT 运算符
 */

func PrintLogic(){
	//如果年龄大于18岁且年龄小于60岁
	age := 30
	if age > 18 && age < 60 {
		fmt.Println("上班族")
	}
}
