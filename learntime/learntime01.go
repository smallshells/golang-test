package learntime

import (
	"fmt"
	"time"
)

// ------------ time包 ----------------
/*
	time包提供了时间的显式和测量用的函数。日历的计算采用公历。
	时间类型：
		time.Time 类型表示时间。可以通过 time.Now() 函数获取当前的时间对象，然后获取时间
		对象的年月日时分秒等信息。
*/

func TimeDome() {
	now := time.Now() //获取当前时间
	fmt.Println("current time : ", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
