package main

import (
	"fmt"
	"time"
)

func testTime1() {
	var t = time.Now() //返回Time构建
	nans := t.UnixNano()
	micros := t.UnixNano() / int64(time.Microsecond)
	mills := t.UnixNano() / int64(time.Millisecond)
	secs := t.UnixNano() / int64(time.Second)
	fmt.Println(t)
	fmt.Println("纳秒：", nans)
	fmt.Println("微秒：", micros)
	fmt.Println("毫秒：", mills)
	fmt.Println("秒：", secs)
}

//时间格式化显示
func testTime2() {
	t := time.Now()
	t1 := t.String()
	fmt.Println(t1) //2019-06-20 22:38:34.8048062 +0800 CST m=+0.005859201
	timeStr := t1[:19]
	fmt.Println(timeStr) //2019-06-20 22:38:34
	fmt.Printf("%d-%d-%d %d:%d;%d:%d %v ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)

	//指定显示格式
	const shortForm = "2006-01-02 15:04:05"
	str := t.Format(shortForm)
	fmt.Println(str)

	//"2019-06-06T15:58:11+08:00"
	fmt.Println(time.Now().Format("2006-01-02T15:04:05+08:00"))
}

//设置时间
func testTime3() {
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	fmt.Println(temp)
}

//从字符串解析出时间
func testTime4() {
	const format = "2006-01-02 15:04:05"
	timeStr := "2018-01-09 20:24:20"
	p, err := time.Parse(format, timeStr)
	if err == nil {
		fmt.Println(p)
	}

	//从Asia/Shanghai时间转CEST时间
	loc, _ := time.LoadLocation("Asia/Shanghai")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)
}

//时间加减运算
func testTime5() {
	t := time.Now()
	until := t.Add(time.Second * time.Duration(5)) //加10秒
	fmt.Println(t)
	fmt.Println(until)
	for time.Now().Before(until) {
		fmt.Println("每 2 打印一次")
		time.Sleep(time.Second * time.Duration(2))
	}
}

//两个时间的差
func testTime6() {
	t1 := time.Now()
	time.Sleep(time.Second * time.Duration(2))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func main() {
	//打印utc的纳秒、微秒、毫秒
	//testTime1()

	//时间格式化显示
	//testTime2()

	//设置时间
	//testTime3()

	//从字符串解析出时间
	//testTime4()

	//时间加减运算
	//testTime5()

	//两个时间的差
	testTime6()
}
