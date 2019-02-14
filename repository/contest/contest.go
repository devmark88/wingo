package contest

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/messages"
	"gitlab.com/mt-api/wingo/model"
)

// Contest : Contest repository
type Contest struct{}

// GetContest : get contest with meta data by ID
func (c *Contest) GetContest(id uint, db *gorm.DB) (*model.Contest, error) {
	var contest model.Contest
	if result := db.Where("id = ?", id).First(&contest); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "contest", "id", id))
	}
	var meta model.ContestMeta
	if result := db.Where("id = ?", contest.ContestMetaID).First(&meta); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "contest meta", "id", id))
	}
	contest.Meta = meta
	return &contest, nil
}

// GetContestByMeta : get contest meta by id
func (c *Contest) GetContestByMeta(metaID uint, db *gorm.DB) (*model.Contest, error) {
	var contest model.Contest
	if result := db.Where("contest_meta_id = ?", metaID).First(&contest); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "contest", "id", metaID))
	}
	return &contest, nil
}

// GetContestOfQuestion : get contest by question id
func (c *Contest) GetContestOfQuestion(id uint, db *gorm.DB) (*model.Contest, error) {
	var contest model.Contest
	var question model.Question
	if qResult := db.Where("id = ?", id).First(&question); qResult.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "question", "id", id))
	}
	if cResult := db.Where("id = ?", question.ContestID).First(&contest); cResult.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "contest", "id", question.ContestID))
	}
	return &contest, nil
}
