package contest

import (
	"fmt"
	"time"

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
		return fmt.Errorf(fmt.Sprintf("We have a contest in this range. ID: %v", contest.ID))
	}
	db.Create(m)
	fmt.Println(m)
	return nil
}
