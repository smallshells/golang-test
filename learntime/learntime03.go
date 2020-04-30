package learntime

import (
	"fmt"
	"time"
)

// ------------ 定时器 ----------------
/*
	使用time.Tick(时间间隔)来设置定时器
*/

func TickDome() {
	ticker := time.Tick(time.Second)
	n := 10
	for i := range ticker {
		fmt.Println(i)  // 每秒执行一次任务
		if n >= 0{
			fmt.Println(n)
			n--
		}else{
			break
		}
	}
}
