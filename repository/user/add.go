package user

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"gitlab.com/mt-api/wingo/q"

	"gitlab.com/mt-api/wingo/messages"

	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
)

// SaveRepository : Save repository
type SaveRepository struct{}

// SaveUserInfo : add user info to database
func (r *SaveRepository) SaveUserInfo(u *model.UserInfo, db *gorm.DB, srv *machinery.Server) error {
	if result := db.Where(model.UserInfo{ID: u.ID}).Attrs(model.UserInfo{Correctors: u.Correctors, Tickets: u.Tickets}).FirstOrCreate(u); result.Error != nil {
		return fmt.Errorf(fmt.Sprintf(messages.GeneralDBError, result.GetErrors()))
	}
	qm := q.QueueManager{}
	return qm.PublishUserInfo(u, srv)
}
