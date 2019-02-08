package helpers

import "time"

func TimeInTehran(t time.Time) time.Time {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		t = t.In(loc)
	}
	return t
}
