package main

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func main() {
	now()
	parseDate()
	addMonth()
}

func now() {
	now := time.Now()
	fmt.Println(now) // JST

	jst, _ := time.LoadLocation("Asia/Tokyo")
	now = time.Now().In(jst)
	fmt.Println(now) // JST

	now = time.Now().UTC()
	fmt.Println(now) // UTC
}

func parseDate() {
	tstr, _ := time.Parse("2006-01", "2021-01")
	fmt.Println(tstr) // 2021-01-01 00:00:00 +0000 UTC

	tstr2, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 15:04:05")
	fmt.Println(tstr2) // 2021-01-02 15:04:05 +0000 UTC
}

func addMonth() {
	tstr2, _ := time.Parse("2006-01-02", "2021-12-02")
	fmt.Println(tstr2) // 2021-12-02 00:00:00 +0000 UTC

	tstr3 := time.Date(tstr2.Year(), tstr2.Month()+1, tstr2.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Println(tstr3) // 2022-01-02 00:00:00 +0000 UTC

	tstr4 := tstr2.AddDate(0, 2, 0)
	fmt.Println(tstr4) // 2022-02-02 00:00:00 +0000 UTC
}
