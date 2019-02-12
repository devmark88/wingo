package user

import (
	"fmt"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

// GetRepository : repository for getter user
type GetRepository struct{}

// GetUserInfo : Get user info from database
func (r *GetRepository) GetUserInfo(u string, db *gorm.DB) (*model.UserInfo, error) {
	var d model.UserInfo
	fmt.Println(u)
	if result := db.Where("id=?", u).Find(&d); result.Error != nil {
		return nil, fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	if len(d.ID) == 0 {
		return nil, fmt.Errorf(messages.ObjectNotFound, "user", "ID", u)
	}
	return &d, nil
}
