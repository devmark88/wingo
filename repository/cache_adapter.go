package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"gitlab.com/mt-api/wingo/helpers"

	"github.com/go-redis/redis"

	"gitlab.com/mt-api/wingo/model"
)

type CacheAdapter struct {
	Connection *redis.Client
}

func (c *CacheAdapter) GetUserInfo(id string) *model.UserInfo {
	return nil
}
func (c *CacheAdapter) ContestMeta() []*model.ContestMeta {
	return nil
}
func (c *CacheAdapter) SetContestMeta(v []*model.ContestMeta) error {
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
