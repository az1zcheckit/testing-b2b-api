package redis

import (
	"context"
	"time"

	rds "github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

var NewRedis = fx.Provide(newRedis)

type IRedis interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) (err error)
	Get(ctx context.Context, key string) (result string, err error)
}

type redis struct {
	rds *rds.Client
}

func newRedis() IRedis {

	rdb := rds.NewClient(&rds.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return redis{rdb}
}

func (r redis) Set(ctx context.Context, key string, value string, expiration time.Duration) (err error) {
	err = r.rds.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return
	}
	return
}

func (r redis) Get(ctx context.Context, key string) (result string, err error) {
	result, err = r.rds.Get(ctx, key).Result()
	if err != nil {
		return
	}
	return
}
