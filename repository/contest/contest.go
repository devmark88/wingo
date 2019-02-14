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
	var questions []model.Question
	if result := db.Where("contest_id = ?", contest.ID).Find(&questions); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.ObjectNotFound, "questions", "contestID", id))
	}
	contest.Meta = meta
	contest.Questions = questions
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

// GetUserTracks : get tracks of user for specific contest
func (c *Contest) GetUserTracks(uID string, cID uint, db *gorm.DB) (*[]model.UserTrack, error) {
	var t []model.UserTrack
	if res := db.Where("user_id = ? AND contest_id = ?", uID, cID); res.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, res.Error))
	}
	return &t, nil
}

// SaveUserTracks : add new track for user into db
func (c *Contest) SaveUserTracks(t *model.UserTrack, db *gorm.DB) error {
	if result := db.Create(t); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	return nil
}
