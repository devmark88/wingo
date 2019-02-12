package contest

import (
	"fmt"
	"time"

	"gitlab.com/mt-api/wingo/helpers"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"

	"github.com/jinzhu/now"
	"gitlab.com/mt-api/wingo/model"
)

// MetaRepository : Metadata repository
type MetaRepository struct{}

// SaveMeta : Add metadata to the database
func (r *MetaRepository) SaveMeta(m *model.ContestMeta, db *gorm.DB) error {
	m.BeginTime = helpers.TimeInTehran(m.BeginTime)
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

// GetTodayMeta : get today contests meta data
// if force is true it will return {limit} number of meta data
// which they will start in the future
func (r *MetaRepository) GetTodayMeta(db *gorm.DB, force bool, limit int) (*[]model.ContestMeta, error) {
	s := helpers.TimeInTehran(now.BeginningOfDay())
	e := helpers.TimeInTehran(now.EndOfDay())
	d := []model.ContestMeta{}
	db.Where("begin_time BETWEEN ? AND ?", s, e).Order("begin_time asc").Find(&d)

	if force == true && len(d) == 0 {
		db.Where("begin_time > ?", s).Order("begin_time asc").Limit(limit).Find(&d)

		if len(d) == 0 {
			db.Order("id desc").Limit(limit).Find(&d)
			reverseMetaArray(d)
		}
	}
	if db.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, db.GetErrors()))
	}
	return &d, nil
}

func reverseMetaArray(m []model.ContestMeta) {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
}
