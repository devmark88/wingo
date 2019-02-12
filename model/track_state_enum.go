package model

import (
	"database/sql/driver"
)

// TrackStateEnum : got_question, post_answer, win, lose
type TrackStateEnum string

const (
	// GotQuestion : got_question
	GotQuestion TrackStateEnum = "got_question"
	// PostAnswer : post_answer
	PostAnswer TrackStateEnum = "post_answer"
	// WinTheGame : win
	WinTheGame TrackStateEnum = "win"
	// LoseTheGame : lose
	LoseTheGame TrackStateEnum = "lose"
)

// Scan : set TrackStateEnum of current model
func (e *TrackStateEnum) Scan(value interface{}) error {
	*e = TrackStateEnum(value.([]byte))
	return nil
}

// Value : get TrackStateEnum of current model
func (e TrackStateEnum) Value() (driver.Value, error) {
	return string(e), nil
}
