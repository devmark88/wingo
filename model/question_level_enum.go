package model

import (
	"database/sql/driver"
)

// QuestionLevelEnum : mindless, easy, medium, hard, curel, unknown from easy to hard
type QuestionLevelEnum string

const (
	// Mindless : mindless is easier that easy
	Mindless QuestionLevelEnum = "mindless"
	// Easy : easy
	Easy QuestionLevelEnum = "easy"
	// Medium : medium
	Medium QuestionLevelEnum = "medium"
	// Hard : hard
	Hard QuestionLevelEnum = "hard"
	// Cruel : cruel
	Cruel QuestionLevelEnum = "cruel"
	// Unknown : unknown (default)
	Unknown QuestionLevelEnum = "unknown"
)

// Scan : set QuestionLevelEnum of current model
func (e *QuestionLevelEnum) Scan(value interface{}) error {
	*e = QuestionLevelEnum(value.([]byte))
	return nil
}

// Value : get QuestionLevelEnum of current model
func (e QuestionLevelEnum) Value() (driver.Value, error) {
	return string(e), nil
}
