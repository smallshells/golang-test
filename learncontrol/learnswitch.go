package learncontrol

import "fmt"

//-------------- switch -------------
/*
	使用switch语句可方便地对大量的值进行条件判断。
	Go语言规定，每个 switch 只能有一个 default 分支。
*/

func PrintSwitch() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	//一个分支可以有多个值，多个 case 值中间使用英文逗号分隔
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough 语法可以执行满足条件的case的下一个case ， 是为了兼容C预演中的case设计的
	//输出结果为a, b
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough //执行下一条语句
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

}
