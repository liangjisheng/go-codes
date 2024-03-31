package jinzhu_now__test

import (
	"github.com/jinzhu/now"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	t.Log(time.Now()) // 2013-11-18 17:51:49.123456789 Mon

	t.Log(now.BeginningOfMinute())  // 2013-11-18 17:51:00 Mon
	t.Log(now.BeginningOfHour())    // 2013-11-18 17:00:00 Mon
	t.Log(now.BeginningOfDay())     // 2013-11-18 00:00:00 Mon
	t.Log(now.BeginningOfWeek())    // 2013-11-17 00:00:00 Sun
	t.Log(now.BeginningOfMonth())   // 2013-11-01 00:00:00 Fri
	t.Log(now.BeginningOfQuarter()) // 2013-10-01 00:00:00 Tue
	t.Log(now.BeginningOfYear())    // 2013-01-01 00:00:00 Tue

	t.Log(now.EndOfMinute())  // 2013-11-18 17:51:59.999999999 Mon
	t.Log(now.EndOfHour())    // 2013-11-18 17:59:59.999999999 Mon
	t.Log(now.EndOfDay())     // 2013-11-18 23:59:59.999999999 Mon
	t.Log(now.EndOfWeek())    // 2013-11-23 23:59:59.999999999 Sat
	t.Log(now.EndOfMonth())   // 2013-11-30 23:59:59.999999999 Sat
	t.Log(now.EndOfQuarter()) // 2013-12-31 23:59:59.999999999 Tue
	t.Log(now.EndOfYear())    // 2013-12-31 23:59:59.999999999 Tue

	now.WeekStartDay = time.Monday // Set Monday as first day, default is Sunday
	t.Log(now.EndOfWeek())         // 2013-11-24 23:59:59.999999999 Sun

	tt := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	t.Log(now.With(tt).EndOfMonth()) // 2013-02-28 23:59:59.999999999 Thu
}

func TestDemo1(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")

	myConfig := &now.Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}

	tt := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.Now().Location()) // // 2013-11-18 17:51:49.123456789 Mon
	t.Log(myConfig.With(tt).BeginningOfWeek())                                  // 2013-11-18 00:00:00 Mon

	tt1, _ := myConfig.Parse("2002-10-12 22:14:01") // 2002-10-12 22:14:01
	t.Log(tt1)
	tt2, err := myConfig.Parse("2002-10-12 22:14") // returns error 'can't parse string as time: 2002-10-12 22:14'
	t.Log(tt2, err)

	t.Log(now.Monday()) // 2013-11-18 00:00:00 Mon
	//now.Monday("17:44")       // 2013-11-18 17:44:00 Mon
	t.Log(now.Sunday()) // 2013-11-24 00:00:00 Sun (Next Sunday)
	//now.Sunday("18:19:24")    // 2013-11-24 18:19:24 Sun (Next Sunday)
	t.Log(now.EndOfSunday()) // 2013-11-24 23:59:59.999999999 Sun (End of next Sunday)

	tt3 := time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.Now().Location()) // 2013-11-24 17:51:49.123456789 Sun
	t.Log(now.With(tt3).Monday())                                                // 2013-11-18 00:00:00 Mon (Last Monday if today is Sunday)
	//now.With(tt3).Monday("17:44")       // 2013-11-18 17:44:00 Mon (Last Monday if today is Sunday)
	t.Log(now.With(tt3).Sunday()) // 2013-11-24 00:00:00 Sun (Beginning Of Today if today is Sunday)
	//now.With(tt3).Sunday("18:19:24")    // 2013-11-24 18:19:24 Sun (Beginning Of Today if today is Sunday)
	t.Log(now.With(tt3).EndOfSunday()) // 2013-11-24 23:59:59.999999999 Sun (End of Today if today is Sunday)
}

func TestDemo2(t *testing.T) {
	t.Log(time.Now()) // 2013-11-18 17:51:49.123456789 Mon

	// Parse(string) (time.Time, error)
	tt, err := now.Parse("2017") // 2017-01-01 00:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("2017-10") // 2017-10-01 00:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("2017-10-13") // 2017-10-13 00:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("1999-12-12 12") // 1999-12-12 12:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("1999-12-12 12:20") // 1999-12-12 12:20:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("1999-12-12 12:20:21") // 1999-12-12 12:20:21, nil
	t.Log(tt, err)
	tt, err = now.Parse("10-13") // 2013-10-13 00:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("12:20") // 2013-11-18 12:20:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("12:20:13") // 2013-11-18 12:20:13, nil
	t.Log(tt, err)
	tt, err = now.Parse("14") // 2013-11-18 14:00:00, nil
	t.Log(tt, err)
	tt, err = now.Parse("99:99") // 2013-11-18 12:20:00, Can't parse string as time: 99:99
	t.Log(tt, err)

	// MustParse must parse string to time or it will panic
	t.Log(now.MustParse("2013-01-13"))       // 2013-01-13 00:00:00
	t.Log(now.MustParse("02-17"))            // 2013-02-17 00:00:00
	t.Log(now.MustParse("2-17"))             // 2013-02-17 00:00:00
	t.Log(now.MustParse("8"))                // 2013-11-18 08:00:00
	t.Log(now.MustParse("2002-10-12 22:14")) // 2002-10-12 22:14:00
	//now.MustParse("99:99")                   // panic: Can't parse string as time: 99:99
}
