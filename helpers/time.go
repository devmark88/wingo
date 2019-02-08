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
