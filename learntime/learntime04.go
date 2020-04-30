package learntime

import (
	"fmt"
	"time"
)

// ------------ 时间格式化 ----------------
/*
	时间类型有一个自带的方法 format 进行格式化，需要注意的是Go语言中格式化时间不是常见的
	Y-m-d H:M:S 而是使用Go的诞生时间2006年1月2号15点04分05秒999毫秒(记忆口诀为2006 1 2 3 4)。也许
	这就是技术人员的浪漫吧。
*/

func FormatDome() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.999"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006-01-02"))
	fmt.Println(now.Format("2006/01/02"))
}
