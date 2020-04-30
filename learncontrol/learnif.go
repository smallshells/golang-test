package learncontrol

import "fmt"

//------------ if-else --------------

func PrintIf() {
	age1 := 30
	if age1 > 18 {
		fmt.Println("你已经成年了，该独立了！")
	} else {
		fmt.Println("你还是个小屁孩")
	}

	//多个判断
	if age1 > 35 {
		fmt.Println("人到中年不得已")
	} else if age1 > 18 {
		fmt.Println("小青年一个")
	} else {
		fmt.Println("好好学习天天向上")
	}

	//age2变量此时只在if条件判断语句中生效
	//if语句执行完，age2变量就被释放了，节省内存空间
	if age2 := 19; age2 > 18 {
		fmt.Println("你已经成年了，该独立了！")
	} else {
		fmt.Println("你还是个小屁孩")
	}
}
