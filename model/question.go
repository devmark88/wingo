package model

// Question : Question Model
type Question struct {
	Base
	Body      string `gorm:"size:100"`
	ContestID uint
	Answers   string
	Level     QuestionLevelEnum
	Order     byte
}
