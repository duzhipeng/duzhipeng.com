package utils

import (
	"bot/config"
	"crypto/tls"
	"github.com/hibiken/asynq"
	"time"
)

var AsynqClient *asynq.Client
var AsynqScheduler *asynq.Scheduler

func InitAsynqClient(c config.Configs) {
	// Redis
	var redis asynq.RedisClientOpt
	if !c.Debug {
		redis = asynq.RedisClientOpt{
			Addr: c.Redis.URI,
			// Omit if no password is required
			Password: c.Redis.Password,
			// Use a dedicated db number for asynq.
			// By default, Redis offers 16 databases (0..15)
			DB: 12,
			// TLS will be negotiated only if this field is set.
			TLSConfig: &tls.Config{},
		}
	} else {
		redis = asynq.RedisClientOpt{
			Addr: c.Redis.URI,
			// Omit if no password is required
			Password: c.Redis.Password,
			// Use a dedicated db number for asynq.
			// By default, Redis offers 16 databases (0..15)
			DB: 12,
			// TLS will be negotiated only if this field is set.
			//TLSConfig: &tls.Config{},
		}
	}
	// Client
	AsynqClient = asynq.NewClient(redis)
	// Scheduler
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	AsynqScheduler = asynq.NewScheduler(redis, &asynq.SchedulerOpts{
		Location: loc,
	})
}

func InitAsynqServer(c config.Configs) *asynq.Server {
	// Redis
	var redis asynq.RedisClientOpt
	if !c.Debug {
		redis = asynq.RedisClientOpt{
			Addr: c.Redis.URI,
			// Omit if no password is required
			Password: c.Redis.Password,
			// Use a dedicated db number for asynq.
			// By default, Redis offers 16 databases (0..15)
			DB: 12,
			// TLS will be negotiated only if this field is set.
			TLSConfig: &tls.Config{},
		}
	} else {
		redis = asynq.RedisClientOpt{
			Addr: c.Redis.URI,
			// Omit if no password is required
			Password: c.Redis.Password,
			// Use a dedicated db number for asynq.
			// By default, Redis offers 16 databases (0..15)
			DB: 12,
			// TLS will be negotiated only if this field is set.
			//TLSConfig: &tls.Config{},
		}
	}
	// Server
	srv := asynq.NewServer(
		redis,
		asynq.Config{Concurrency: 10},
	)
	return srv
}
