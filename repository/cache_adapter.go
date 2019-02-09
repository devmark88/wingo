package repository

import (
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

func getDateForKey(t time.Time) string {
	t = helpers.TimeInTehran(t)
	return fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
}
