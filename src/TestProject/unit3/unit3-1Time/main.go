package main

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func main() {
	// 獲取當前時間
	t0 := time.Now()
	fmt.Println("t0 =", t0.Unix())

	// 程式根據給定時間暫停
	time.Sleep(50 * time.Millisecond)

	// 時間相減
	t1 := time.Now()
	subTime := t1.Sub(t0).Milliseconds()
	fmt.Println("t1 - t0 =", subTime)

	// 從某時間到當前時間間隔
	since := time.Since(t0).Milliseconds()
	fmt.Println("time of now since t0 =", since)

	// 時間相加
	d := time.Duration(2 * time.Second)
	t2 := t0.Add(d)
	fmt.Println("t0 + d =", t2.Unix())

	// 時間轉字串
	fmt.Println("yyyy-mm-dd =", t0.Format(DATE))
	s := t0.Format(TIME)
	fmt.Println("yyyy-mm-dd h:m:s =", s)

	// 字串轉時間
	loc, _ := time.LoadLocation("Asia/Taipei")
	t3, _ := time.ParseInLocation(TIME, s, loc)
	fmt.Println("time parse =", t3.Unix())
}
