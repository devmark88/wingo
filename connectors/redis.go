package connectors

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func CreateRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.address"),
		PoolSize:     viper.GetInt("redis.pool"),
		DialTimeout:  viper.GetDuration("redis.timeout.dial"),
		ReadTimeout:  viper.GetDuration("redis.timeout.read"),
		WriteTimeout: viper.GetDuration("redis.timeout.write"),
		PoolTimeout:  viper.GetDuration("redis.timeout.pool"),
		IdleTimeout:  viper.GetDuration("redis.timeout.idle"),
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	return client
}
