package repository

import (
	"gitlab.com/mt-api/wingo/model"
)

type CacheAdapter struct{}

func (c *CacheAdapter) GetUserInfo(id string) *model.UserInfo {
	return nil
}
