package model

import (
	"database/sql/driver"
)

type QuestionLevelEnum string

const (
	Mindless QuestionLevelEnum = "mindless"
	Easy     QuestionLevelEnum = "easy"
	Medium   QuestionLevelEnum = "medium"
	Hard     QuestionLevelEnum = "hard"
	Cruel    QuestionLevelEnum = "cruel"
	Unknown  QuestionLevelEnum = "unknown"
)

func (e *QuestionLevelEnum) Scan(value interface{}) error {
	*e = QuestionLevelEnum(value.([]byte))
	return nil
}

func (e QuestionLevelEnum) Value() (driver.Value, error) {
	return string(e), nil
}
