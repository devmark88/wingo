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

func (r *MetaRepository) SaveContest(m *model.Contest, db *gorm.DB) error {
	x := model.Contest{}
	db.Table("contests").Where("contest_meta_id = ?", m.ContestMetaID).Find(&x)
	if x.ID > 0 {
		return fmt.Errorf(fmt.Sprintf(messages.META_HAS_CONTEST, x.ID))
	}
	tx := db.Begin()
	ct := model.Contest{
		ContestMetaID:         m.ContestMetaID,
		CorrectAnswersIndices: m.CorrectAnswersIndices,
	}
	if result := db.Create(&ct); result.Error != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
	}
	for _, q := range m.Questions {
		question := model.Question{
			Answers:   q.Answers,
			Body:      q.Body,
			ContestID: ct.ID,
			Level:     q.Level,
			Order:     q.Order,
		}
		if result := db.Create(&question); result.Error != nil {
			tx.Rollback()
			return fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
		}
	}
	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
	}
	return nil
}
