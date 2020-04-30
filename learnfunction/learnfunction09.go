package learnfunction

import "fmt"

//--------------- 递归 -------------------
/*
	就是自己调用自己
	递归要有一个明确的退出条件
	适合处理那种问题相同、问题的规模越来越小的规模
*/

//阶乘
func factorial(n uint64) (result uint64) {
	if n > 1 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

//n个台阶，一次可以走1步，也可以走2步，有多少种走法
func step(n int) int {
	if 1 == n {
		return 1
	}
	if 2 == n {
		return 2
	}
	return step(n-1) + step(n-2)
}

func PrintFunction09() {
	fmt.Println(factorial(10))
	fmt.Println(step(10))
}
