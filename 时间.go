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
	fmt.Println("Time ", t)
	fmt.Println("纳秒：", nans)
	fmt.Println("微秒：", micros)
	fmt.Println("毫秒：", mills)
	fmt.Println("秒：", secs)
}
