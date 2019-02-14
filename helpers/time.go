package helpers

import (
	"time"
)

// TimeInTehran : Change timezone of specific time to GMT+3:30
func TimeInTehran(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Tehran")
	t = t.In(loc)
	return t
}

// GetTime : Get current time (hour, minute, seconds)
func GetTime(t time.Time) (hour int, minute int, second int) {
	layout := "2006-01-02 15:04:05"
	timeStampString := t.Format(layout)
	timeStamp, err := time.Parse(layout, timeStampString)
	if err != nil {
		return 0, 0, 0
	}
	return timeStamp.Clock()
}

// GetDate : return yearn, month, and day
func GetDate(t time.Time) (year int, month int, day int) {
	return t.Year(), int(t.Month()), t.Day()
}

// IsPast : check wheter input time is past
// Convert input time to GMT + 3:30
func IsPast(t time.Time) bool {
	d := TimeInTehran(t)
	n := TimeInTehran(time.Now())
	return n.Sub(d) > 0
}

// DateWithinRange : check wheter t is in the range
// Convert input to GMT + 3:30
func DateWithinRange(t time.Time, min time.Time, max time.Time) bool {
	min = TimeInTehran(min)
	max = TimeInTehran(max)
	t = TimeInTehran(t)

	if t.Before(max) && t.After(min) {
		return true
	}
	return false
}

// IsFuture : check wheter input time is past
// Convert input time to GMT + 3:30
func IsFuture(t time.Time) bool {
	d := TimeInTehran(t)
	n := TimeInTehran(time.Now())
	return n.Sub(d) < 0
}
