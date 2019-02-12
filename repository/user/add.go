package user

import (
	"fmt"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

// SaveRepository : Save repository
type SaveRepository struct{}

// SaveUserInfo : add user info to database
func (r *SaveRepository) SaveUserInfo(u *model.UserInfo, db *gorm.DB) error {
	if result := db.Where(model.UserInfo{ID: u.ID}).Attrs(model.UserInfo{Correctors: u.Correctors, Tickets: u.Tickets}).FirstOrCreate(u); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	return nil
}
