package repository

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/logger"
	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/repository/contest"
	"gitlab.com/mt-api/wingo/repository/user"
)

type Connections struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (cn *Connections) AddMeta(m *model.ContestMeta) error {
	r := contest.MetaRepository{}
	err := r.SaveMeta(m, cn.DB)
	if err == nil {
		c := CacheAdapter{Connection: cn.Redis}
		er := c.InvalidateContestMeta()
		if er != nil {
			logger.Error(fmt.Errorf("error while invalidating meta into redis: %v", er))
		}
	}
	return err
}

func (cn *Connections) GetMeta(force bool) (error, []*model.ContestMeta) {
	r := contest.MetaRepository{}
	err, t := r.GetTodayMeta(cn.DB, force, 3)
	if err == nil {
		c := CacheAdapter{Connection: cn.Redis}
		er := c.SetContestMeta(t)
		if er != nil {
			logger.Error(fmt.Errorf("error while saving meta into redis: %v", er))
		}
	}
	return err, t
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
