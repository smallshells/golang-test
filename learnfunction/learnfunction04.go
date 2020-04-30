package learnfunction

import "fmt"

// ---------------- 函数类型作为参数和返回值 ---------------

func f1() {
	fmt.Println("Hello 俊林")
}

func f2() int {
	return 10
}

func f3(x int) int {
	return x + 6
}

//函数作为参数类型
func f4(x func() int) {
	ret := x()
	fmt.Println(ret)
}

//函数作为返回值
func f5() func(int, int) int {
	ret := func(x, y int) int {
		return x + y
	}
	return ret
}

//参数和返回值都是函数类型
func f6(x func(int) int) func(int, int) int {
	a := x(4)
	ret := func(x, y int) int {
		return x + y + a
	}
	return ret
}

func PrintFunction04() {
	a1 := f1
	a2 := f2
	a3 := f3
	fmt.Printf("The type of 'a1' is %T\n", a1)
	fmt.Printf("The type of 'a2' is %T\n", a2)
	fmt.Printf("The type of 'a3' is %T\n", a3)

	f4(f2)
	f4(a2)
	a5 := f5()
	fmt.Println(a5(4, 5))

	a6 := f6(f3)
	fmt.Println(a6(5,6))
}
