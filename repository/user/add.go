package user

import (
	"fmt"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

type UserSaveRepository struct{}

func (r *UserSaveRepository) SaveUserInfo(u *model.UserInfo, db *gorm.DB) error {
	if result := db.Create(u); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GENERAL_DB_ERROR, result.GetErrors()))
	}
	return nil
}
