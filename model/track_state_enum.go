package model

import (
	"database/sql/driver"
)

type TrackStateEnum string

const (
	GotQuestion TrackStateEnum = "got_question"
	PostAnswer  TrackStateEnum = "post_answer"
	WinTheGame  TrackStateEnum = "win"
	LoseTheGame TrackStateEnum = "lose"
)

func (e *TrackStateEnum) Scan(value interface{}) error {
	*e = TrackStateEnum(value.([]byte))
	return nil
}

func (e TrackStateEnum) Value() (driver.Value, error) {
	return string(e), nil
}
