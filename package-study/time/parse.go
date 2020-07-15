package main

import (
	"fmt"
	"strconv"
	"time"
)

func parse() {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Jun 21, 2017 at 0:00am (PST)")
	fmt.Println(t)

	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2017-Jun-21")
	fmt.Println(t)

	t, _ = time.Parse("01/02/2006", "06/21/2017")
	fmt.Println(t)
	fmt.Println(t.Unix())

	i, err := strconv.ParseInt("1498003200", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)

	var timestamp int64 = 1498003200
	tm2 := time.Unix(timestamp, 0)
	fmt.Println(tm2.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(tm2.Format("02/01/2006 15:04:05 PM"))

	t, err = time.Parse("2006-Jan-02 15:04:05 UTC", "2020-May-05 02:42:22.000000000 UTC")
	if err != nil {
		fmt.Println("parse err:", err)
		return
	}
	fmt.Printf("time: %+v", t)
}
