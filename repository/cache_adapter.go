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

// CacheAdapter : cache adapter
type CacheAdapter struct {
	Connection *redis.Client
}

// GetUserInfo : get user info cache
// with key user:{{userID}}
func (c *CacheAdapter) GetUserInfo(id string) *model.UserInfo {
	i := fmt.Sprintf("user:%s", id)
	var u model.UserInfo
	b, e := c.Connection.Get(i).Bytes()
	if e != nil {
		return nil
	}
	json.Unmarshal(b, &u)
	return &u
}

// SetUserInfo : set user info cache
// with key user:{{userID}}
func (c *CacheAdapter) SetUserInfo(u model.UserInfo) error {
	serialized, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("error in marshal to json : %v", err)
	}
	return c.Connection.Set(fmt.Sprintf("user:%s", u.ID), serialized, 0).Err()
}

// InvalidateUserInfo : invalidate user info cache
// with key user:{{userID}}
func (c *CacheAdapter) InvalidateUserInfo(id string) error {
	return c.Connection.Del(fmt.Sprintf("user:%s", id)).Err()
}

// GetContestMeta : get contest meta from cache
// with key meta:contest:{{year}}-{{month}}-{{day}}
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

// SetContestMeta : set meta contest cache
// with key meta:contest:{{year}}-{{month}}-{{day}}
func (c *CacheAdapter) SetContestMeta(v *[]model.ContestMeta) error {
	serialized, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error in marshal to json : %v", err)
	}
	return c.Connection.Set(fmt.Sprintf("meta:contest:%s", getDateForKey(time.Now())), serialized, time.Hour*24).Err()
}

// InvalidateContestMeta : invalidate meta contest cache for today
// with key meta:contest:{{year}}-{{month}}-{{day}}
func (c *CacheAdapter) InvalidateContestMeta() error {
	key := fmt.Sprintf("meta:contest:%s", getDateForKey(time.Now()))
	return c.Connection.Del(key).Err()
}

// GetUserTrack : get user track from redis
// with key user:{{userID}}:contest:{{contestID}}:track
func (c *CacheAdapter) GetUserTrack(cid int, uid string) *[]model.UserTrack {
	k := fmt.Sprintf("user:%s:contest:%v:track", uid, cid)
	var cm []model.UserTrack
	b, e := c.Connection.Get(k).Bytes()
	if e != nil {
		return nil
	}
	json.Unmarshal(b, &cm)
	return &cm
}

// SetUserTrack : add new item to user track
// with key user:{{userID}}:contest:{{contestID}}:track
func (c *CacheAdapter) SetUserTrack(v *model.UserTrack) error {
	k := fmt.Sprintf("user:%s:contest:%v:track", v.UserID, v.ContestID)
	serialized, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error in marshal to json: %v", err)
	}
	val := make(map[string]interface{})
	val[fmt.Sprintf("%v", v.QuestionIndex)] = serialized
	return c.Connection.HMSet(k, val).Err()
}

func getDateForKey(t time.Time) string {
	t = helpers.TimeInTehran(t)
	return fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
}
