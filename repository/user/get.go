package user

import (
	"fmt"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

// UserGetRepository : repository for getter user
type UserGetRepository struct{}

// GetUserInfo : Get user info from database
func (r *UserGetRepository) GetUserInfo(u string, db *gorm.DB) (*model.UserInfo, error) {
	var d model.UserInfo
	if result := db.Where("id=?", u).Find(&d); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
	}
	if len(d.ID) == 0 {
		return nil, fmt.Errorf(messages.NOT_FOUND, "user", "ID", u)
	}
	return &d, nil
}
