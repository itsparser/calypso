package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/workfoxes/calypso/pkg/config"
	"github.com/workfoxes/calypso/pkg/constant"
	"github.com/workfoxes/calypso/pkg/errors"
	"github.com/workfoxes/calypso/pkg/log"
	"sync"
	"time"
)

var r = new(Redis)

type Redis struct {
	once     sync.Once
	_default *redis.Client
}

func Init() *Redis {
	var err error = nil
	r.once.Do(func() {
		r._default = redis.NewClient(&redis.Options{
			Addr:     config.GetValue(constant.RedisHost),
			Password: config.GetValue(constant.RedisPassword),
			DB:       0,
		})
		ctx := context.Background()
		if err = r._default.Ping(ctx).Err(); err != nil {
			log.Panic(errors.RedisUnreachable)
		}
	})
	return r
}

// Get : will help used to get the value from redis.
func Get(ctx context.Context, key string) *redis.StringCmd {
	return Init()._default.Get(ctx, key)
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return Init()._default.Set(ctx, key, value, expiration)
}

func Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return Init()._default.Subscribe(ctx, channel)
}
