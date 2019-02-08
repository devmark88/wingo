package user

import (
	"fmt"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

type UserGetRepository struct{}

func (r *UserGetRepository) GetUserInfo(u string, db *gorm.DB) (error, *model.UserInfo) {
	var d model.UserInfo
	db.Where("ID=?", u).Find(&d)
	if len(d.ID) == 0 {
		return fmt.Errorf(messages.NOT_FOUND, "user", "ID", u), nil
	}
	return nil, &d
}
