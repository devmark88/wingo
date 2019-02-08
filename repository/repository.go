package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/repository/contest"
	"gitlab.com/mt-api/wingo/repository/user"
)

type Connections struct {
	DB *gorm.DB
}

func (cn *Connections) AddMeta(m *model.ContestMeta) error {
	r := contest.MetaRepository{}
	return r.SaveMeta(m, cn.DB)
}

func (cn *Connections) GetMeta(force bool) (error, []*model.ContestMeta) {
	r := contest.MetaRepository{}
	return r.GetTodayMeta(cn.DB, force, 3)
}
func (cn *Connections) AddContest(m *model.Contest) error {
	r := contest.QuestionRepository{}
	return r.SaveContest(m, cn.DB)
}
func (cn *Connections) GetUserInfo(id string) (error, *model.UserInfo) {
	r := user.UserGetRepository{}
	return r.GetUserInfo(id, cn.DB)
}
func (cn *Connections) AddUserInfo(u *model.UserInfo) error {
	r := user.UserSaveRepository{}
	return r.SaveUserInfo(u, cn.DB)
}
