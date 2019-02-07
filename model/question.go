package model

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Question struct {
	Base
	Body      string `gorm:"size:100"`
	ContestID uint
	Answers   postgres.Jsonb
	Level     QuestionLevelEnum
}
