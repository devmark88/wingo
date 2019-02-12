package contest

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"

	"gitlab.com/mt-api/wingo/logger"

	"gitlab.com/mt-api/wingo/q"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/messages"
	"gitlab.com/mt-api/wingo/model"
)

//QuestionRepository ...
type QuestionRepository struct{}

//SaveContest : Add questions with answers to corresponding contest
func (r *QuestionRepository) SaveContest(m *model.Contest, db *gorm.DB, srv *machinery.Server) error {
	x := model.Contest{}
	db.Table("contests").Where("contest_meta_id = ?", m.ContestMetaID).Find(&x)
	if x.ID > 0 {
		return fmt.Errorf(fmt.Sprintf(messages.MetaHasContest, x.ID))
	}
	tx := db.Begin()
	ct := model.Contest{
		ContestMetaID:         m.ContestMetaID,
		CorrectAnswersIndices: m.CorrectAnswersIndices,
	}
	if result := db.Create(&ct); result.Error != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
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
			return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
		}
	}
	var meta model.ContestMeta
	db.Where("id = ?", m.ContestMetaID).First(&meta)
	m.Meta = meta
	g := q.Question{}
	e := g.PublishAll(*m, srv)
	if e != nil {
		logger.Error(e)
		logger.Debug("rolling back the transaction")
		tx.Rollback()
		logger.Debug("transaction rolled back")
		return e
	}
	if result := tx.Commit(); result.Error != nil {
		tx.Rollback()
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	return nil
}
