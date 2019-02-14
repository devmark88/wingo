package user

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/messages"
	"gitlab.com/mt-api/wingo/model"
)

type TrackRepository struct{}

// GetUserTracks : get tracks of user for specific contest
func (c *TrackRepository) GetUserTracks(uID string, cID uint, db *gorm.DB) (*[]model.UserTrack, error) {
	var t []model.UserTrack
	if res := db.Where("user_id = ? AND contest_id = ?", uID, cID); res.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, res.Error))
	}
	return &t, nil
}

// SaveUserTracks : add new track for user into db
func (c *TrackRepository) SaveUserTracks(t *model.UserTrack, db *gorm.DB) error {
	if result := db.Create(t); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	return nil
}
