package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"gitlab.com/mt-api/wingo/logger"

	"gitlab.com/mt-api/wingo/helpers"

	"github.com/go-redis/redis"

	"gitlab.com/mt-api/wingo/model"
)

type CacheAdapter struct {
	Connection *redis.Client
}

func (c *CacheAdapter) GetUserInfo(id string) *model.UserInfo {
	i := fmt.Sprintf("user:%s", id)
	var u model.UserInfo
	b, e := c.Connection.Get(i).Bytes()
	if e != nil {
		logger.Error(fmt.Errorf("error while getting user info from redis UID:%s => %v", id, e))
		return nil
	}
	json.Unmarshal(b, &u)
	return &u
}
func (c *CacheAdapter) SetUserInfo(u model.UserInfo) error {
	serialized, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("error in marshal to json : %v", err)
	}
	return c.Connection.Set(fmt.Sprintf("user:%s", u.ID), serialized, 0).Err()
}
func (c *CacheAdapter) InvalidateUserInfo(id string) error {
	return c.Connection.Del(fmt.Sprintf("user:%s", id)).Err()
}

func (c *CacheAdapter) GetContestMeta() *[]model.ContestMeta {
	k := fmt.Sprintf("meta:contest:%s", getDateForKey(time.Now()))
	var cm []model.ContestMeta
	b, e := c.Connection.Get(k).Bytes()
	if e != nil {
		logger.Error(fmt.Errorf("error while getting contest meta from redis Key:%s", k))
		return nil
	}
	json.Unmarshal(b, &cm)
	return &cm
}
func (c *CacheAdapter) SetContestMeta(v *[]model.ContestMeta) error {
	serialized, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error in marshal to json : %v", err)
	}
	return c.Connection.Set(fmt.Sprintf("meta:contest:%s", getDateForKey(time.Now())), serialized, time.Hour*24).Err()
}
func (c *CacheAdapter) InvalidateContestMeta() error {
	key := fmt.Sprintf("meta:contest:%s", getDateForKey(time.Now()))
	return c.Connection.Del(key).Err()
}

func getDateForKey(t time.Time) string {
	t = helpers.TimeInTehran(t)
	return fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
}
