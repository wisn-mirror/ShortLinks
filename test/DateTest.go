package main

import (
	"fmt"
	"time"
)

func main() {
	//2019-10-31 13:48:23.514228 +0800 CST m=+0.000275958
	//fmt.Println(time.Now())
	t:=time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d 星期%d,一年第%d天 时间戳%d 当前秒时间戳%d 当前毫秒时间戳%d 当前纳秒时间戳%d,%d\n",
		t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second(),t.Weekday(),t.YearDay(),t.Unix(),t.UnixNano()/1e9,t.UnixNano()/1e6,t.UnixNano(),t.Nanosecond())
	fmt.Println(t.Format("时间：15:04:05"))
	fmt.Println(t.Format("t 日期：2006-01-02 时间：15:04:05"))
	t2 := t.AddDate(10, 3, 3)
	fmt.Println(t2.Format("t2 日期：2006-01-02 时间：15:04:05"))
	t3 := t2.AddDate(-2, -2, -3)
	fmt.Println(t3.Format("t3 日期：2006-01-02 时间：15:04:05"))
	fmt.Println(t3.After(t2))
}