package main

import (
	"fmt"
	"time"
)

func main() {
	var t = time.Now()
	nans := t.UnixNano()
	micros := t.UnixNano() / int64(time.Microsecond)
	mills := t.UnixNano() / int64(time.Millisecond)
	secs := t.UnixNano() / int64(time.Second)
	fmt.Println(t)
	fmt.Println("纳秒：", nans)
	fmt.Println("微秒：", micros)
	fmt.Println("毫秒：", mills)
	fmt.Println("秒：", secs)

	//2019-06-06T15:58:11+08:00
	{
		fmt.Println("")
		nowTime := time.Now()
		t1 := nowTime.String()
		timeStr := t1[:19]
		fmt.Println(timeStr)
	}

	{
		fmt.Println("")
		const shortForm = "2006-01-01 15:04:05"
		t := time.Now()
		temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		str := temp.Format(shortForm)
		fmt.Println(str)
	}

	GetTimeFromStr()
	GetTimeFromStrLoc()
}

func GetTimeFromStr() {
	fmt.Println("\n GetTimeFromStr")
	const format = "2006-01-02 15:04:05"
	timeStr := "2018-01-09 20:24:20"
	p, err := time.Parse(format, timeStr)
	if err == nil {
		fmt.Println(p)
	}
}

//带时区匹配，匹配当前时区的时间
func GetTimeFromStrLoc() {
	fmt.Println("\n GetTimeFromStrLoc")

	loc, _ := time.LoadLocation("Asia/Shanghai")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)
	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	//"2019-06-06T15:58:11+08:00"
	fmt.Println(time.Now().Format("2006-01-02T15:04:05+08:00"))
}
