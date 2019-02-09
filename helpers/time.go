package helpers

import (
	"time"
)

func TimeInTehran(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Tehran")
	t = t.In(loc)
	return t
}
func GetTime(t time.Time) (hour int, minute int, second int) {
	layout := "2006-01-02 15:04:05"
	timeStampString := t.Format(layout)
	timeStamp, err := time.Parse(layout, timeStampString)
	if err != nil {
		return 0, 0, 0
	}
	return timeStamp.Clock()
}
func GetDate(t time.Time) (year int, month int, day int) {
	return t.Year(), int(t.Month()), t.Day()
}
func IsPast(t time.Time) bool {
	d := TimeInTehran(t)
	n := TimeInTehran(time.Now())
	return n.Sub(d) > 0
}

func IsFuture(t time.Time) bool {
	d := TimeInTehran(t)
	n := TimeInTehran(time.Now())
	return n.Sub(d) < 0
}
