package learncontrol

import "fmt"

//------------ for --------------

func PrintFor() {
	//基本格式
	for i1 := 0; i1 < 10; i1++ {
		fmt.Printf("%d ", i1)
	}
	fmt.Println()

	//变种1
	var i2 = 5
	for ; i2 < 10; i2++ {
		fmt.Printf("%d ", i2)
	}
	fmt.Println()

	//变种2
	var i3 = 2
	for ; i3 < 10; {
		fmt.Printf("%d ", i3)
		i3++
	}
	fmt.Println()

	//变种4
	var i4 = 6
	for i4 < 10 {
		fmt.Printf("%d ", i4)
		i4++
	}
	fmt.Println()

	//无限循环
	//for {
	//	fmt.Println("无限循环中……")
	//}

	//for range 循环
	s := "Hello拉萨"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	//break 跳出循环即结束for循环
	for i:= 0; i < 9; i++{
		if i == 5 {
			break
		}
		fmt.Printf("%d ",i)
	}

	//continue 跳过此次循环，继续下一次循环
	for i:= 0; i < 9; i++{
		if i == 5 {
			continue
		}
		fmt.Printf("%d ",i)
	}
}

//打印九九乘法表
func PrintTable() {
	var i, j int
	for i = 1; i < 10; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d × %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

//打印100以内的素数
func PrintPrimeNumber() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}
