package cache

import (
	"ChargPiles/config"
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitRedis() {
	rConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		//Username: rConfig.RedisUsername,
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})

	RedisClient = client
}
