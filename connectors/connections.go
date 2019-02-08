package connectors

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Connections struct {
	Database *gorm.DB
	Cache    *redis.Client
}
