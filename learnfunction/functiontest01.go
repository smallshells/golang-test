package learnfunction

import "fmt"

/*
你有50个金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth.
分配规则如下：
a.名字中每包含1给'e'或'E'分1枚金币
b.名字中每包含1给'i'或'I'分2枚金币
c.名字中每包含1给'o'或'O'分3枚金币
d.名字中每包含1给'u'或'U'分4枚金币
下一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现'dispatchcion'函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCion() int {
	for _, name := range users {
		distribution[name] = 0
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	fmt.Println(distribution)

	return coins
}

func FunctionTest01() {
	left := dispatchCion()
	fmt.Println("剩下：", left)
}
