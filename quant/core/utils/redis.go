package utils

import (
	"core/config"
	"crypto/tls"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(c config.Configs) {
	if !c.Debug {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     c.Redis.URI,
			Password: c.Redis.Password, // no password set
			DB:       c.Redis.DB,       // use default DB
			// TLS will be negotiated only if this field is set.
			TLSConfig: &tls.Config{},
		})
	} else {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     c.Redis.URI,
			Password: c.Redis.Password, // no password set
			DB:       c.Redis.DB,       // use default DB
		})
	}
}
