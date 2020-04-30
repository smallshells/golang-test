package learntime

import (
	"fmt"
	"time"
)

// ------------ 时间戳 ----------------
/*
	时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
*/

func TimeStampDemo01() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println("current timestamp1 :", timestamp1)
	fmt.Println("current timestamp2 :", timestamp2)
	timestampDemo02(timestamp1 - 1000) //当前时间1000秒之前
}

func timestampDemo02(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
