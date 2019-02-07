package contest

import (
	"fmt"
	"time"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"

	"gitlab.com/mt-api/wingo/model"
)

type MetaRepository struct{}

func (r *MetaRepository) SaveMeta(m *model.ContestMeta, db *gorm.DB) error {
	s := m.BeginTime
	e := m.BeginTime.Add(time.Second * time.Duration(m.Duration))
	var contest model.ContestMeta
	db.Where("begin_time BETWEEN ? AND ?", s, e).First(&contest)
	if contest.ID > 0 {
		return fmt.Errorf(fmt.Sprintf(messages.INVALID_CONTEST_TIME, contest.ID))
	}

	if result := db.Create(m); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
	}
	return nil
}
