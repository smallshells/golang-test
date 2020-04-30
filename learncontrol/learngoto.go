package learncontrol

import "fmt"

//--------------- goto -------------------
/*
	goto 语句通过标签进行代码间的无条件跳转。
	goto 语句可以在快速跳出循环、避免重复退出上有一定的帮助。
	Go预演中使用 goto 语句能简化一些代码的实现过程。
*/

func PrintGoto() {
	//双层嵌套的 for 循环要退出时
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		if breakFlag {
			break
		}
	}

	//使用 goto 语句能简化代码
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
breakTag:
	fmt.Println("结束for循环")
}
