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

// GetTodayContestsMeta : get contest meta from cache
// with key meta:contest:{{year}}-{{month}}-{{day}}
func (c *CacheAdapter) GetTodayContestsMeta() *[]model.ContestMeta {
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
func (c *CacheAdapter) GetUserTrack(cid uint, uid string) *[]model.UserTrack {
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

// SetUserTracks : replace user tracks if exists otherwise add it
// with key user:{{userID}}:contest:{{contestID}}:track
func (c *CacheAdapter) SetUserTracks(v *[]model.UserTrack) error {
	val := make(map[string]interface{})
	for _, q := range *v {
		k := fmt.Sprintf("user:%s:contest:%v:track", q.UserID, q.ContestID)
		serialized, err := json.Marshal(q)
		if err != nil {
			return fmt.Errorf("error in marshal to json: %v", err)
		}
		val[fmt.Sprintf("%v", q.QuestionIndex)] = serialized
		err = c.Connection.HMSet(k, val).Err()
	}
	
	return nil
}

// GetContestMetabyID : get contest meta by id
// with key contest:meta:{{id}}
func (c *CacheAdapter) GetContestMetabyID(id uint) *model.ContestMeta {
	var ct model.ContestMeta
	k := fmt.Sprintf("contest:meta:%v", id)
	b, e := c.Connection.Get(k).Bytes()
	if e != nil {
		return nil
	}
	json.Unmarshal(b, &ct)
	return &ct
}

// InvalidateContestMetaByID : invalidate contest meta cache by id
// with key contest:meta:{{id}}
func (c *CacheAdapter) InvalidateContestMetaByID(id uint) error {
	key := fmt.Sprintf("contest:meta:%v", id)
	return c.Connection.Del(key).Err()
}

// SetContestMetabyID : get contest by id
// with key contest:meta:{{id}}
func (c *CacheAdapter) SetContestMetabyID(v model.ContestMeta) error {
	k := fmt.Sprintf("contest:meta:%v", v.ID)
	serialized, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error in marshal to json: %v", err)
	}
	return c.Connection.Set(k, serialized, 0).Err()
}

// GetContest : get contest by id
// get contest by id with key contest:{{id}}
func (c *CacheAdapter) GetContest(id uint) *model.Contest {
	var ct model.Contest
	k := fmt.Sprintf("contest:%v", id)
	b, e := c.Connection.Get(k).Bytes()
	if e != nil {
		return nil
	}
	json.Unmarshal(b, &ct)
	return &ct
}

// SetContest : get contest by id
// with key contest:{{id}}
func (c *CacheAdapter) SetContest(v model.Contest) error {
	k := fmt.Sprintf("contest:%v", v.ID)
	serialized, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("error in marshal to json: %v", err)
	}
	return c.Connection.Set(k, serialized, 0).Err()
}

// InvalidateContest : invalidate contest cache by id
// with key contest:{{id}}
func (c *CacheAdapter) InvalidateContest(id uint) error {
	key := fmt.Sprintf("contest:%v", id)
	return c.Connection.Del(key).Err()
}

func getDateForKey(t time.Time) string {
	t = helpers.TimeInTehran(t)
	return fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
}
