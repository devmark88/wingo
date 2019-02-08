package helpers

import (
	"time"
)

func TimeInTehran(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Tehran")
	t = t.In(loc)
	return t
}
