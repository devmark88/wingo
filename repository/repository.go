package repository

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gitlab.com/mt-api/wingo/logger"
	"gitlab.com/mt-api/wingo/model"
	"gitlab.com/mt-api/wingo/repository/contest"
	"gitlab.com/mt-api/wingo/repository/user"
)

// Connections : connections for the repository
type Connections struct {
	DB    *gorm.DB
	Redis *redis.Client
	Queue *machinery.Server
}

// AddMeta : add meta data to the database
// it make related cache invalid
func (cn *Connections) AddMeta(m *model.ContestMeta) error {
	r := contest.MetaRepository{}
	err := r.SaveMeta(m, cn.DB)
	if err == nil {
		c := CacheAdapter{Connection: cn.Redis}
		er := c.InvalidateContestMeta()
		if er != nil {
			logger.Error(fmt.Errorf("error while invalidating contest meta cache: %v", er))
		}
	}
	return err
}

// GetMeta : Get meta data of today
// it get data from cache if can found any otherwise it get data from database
// and put it to the cache
func (cn *Connections) GetMeta(force bool) (*[]model.ContestMeta, error) {
	r := contest.MetaRepository{}
	c := CacheAdapter{Connection: cn.Redis}
	t := c.GetContestMeta()
	if t == nil {
		t, err := r.GetTodayMeta(cn.DB, force, 3)
		if err == nil {
			er := c.SetContestMeta(t)
			if er != nil {
				logger.Error(fmt.Errorf("error while saving meta into redis: %v", er))
			}
		}
		return t, err
	}
	return t, nil
}

// AddContest : AddContest to the database
func (cn *Connections) AddContest(m *model.Contest) error {
	r := contest.QuestionRepository{}
	return r.SaveContest(m, cn.DB, cn.Queue)
}

// GetUserInfo : get user info by id
func (cn *Connections) GetUserInfo(id string) (*model.UserInfo, error) {
	r := user.GetRepository{}
	c := CacheAdapter{Connection: cn.Redis}
	u := c.GetUserInfo(id)
	if u == nil {
		u, err := r.GetUserInfo(id, cn.DB)
		if err == nil {
			er := c.SetUserInfo(*u)
			if er != nil {
				logger.Error(fmt.Errorf("error while saving user info into redis user:%v => err: %v", u, er))
			}
		}
		return u, err
	}
	return u, nil
}

// AddUserInfo : add user info to the database
func (cn *Connections) AddUserInfo(u *model.UserInfo) error {
	r := user.SaveRepository{}
	err := r.SaveUserInfo(u, cn.DB)
	if err != nil {
		c := CacheAdapter{Connection: cn.Redis}
		e := c.InvalidateUserInfo(u.ID)
		if e != nil {
			logger.Error(fmt.Errorf("error while invalidating user info cache: %v", e))
		}
	}
	return err
}
