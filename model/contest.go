package model

import (
	"database/sql"

	"github.com/lib/pq"
)

type Contest struct {
	Base
	ContestMetaID         sql.NullInt64
	Meta                  ContestMeta
	Questions             []Question
	CorrectAnswersIndices pq.StringArray `gorm:"type:varchar(100)[]"`
}
