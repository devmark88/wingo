package model

// Contest : Contest Model
type Contest struct {
	Base
	ContestMetaID         uint64
	Meta                  ContestMeta
	Questions             []Question
	CorrectAnswersIndices string
}
