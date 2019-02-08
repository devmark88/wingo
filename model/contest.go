package model

type Contest struct {
	Base
	ContestMetaID         uint64
	Meta                  ContestMeta
	Questions             []Question
	CorrectAnswersIndices string
}
