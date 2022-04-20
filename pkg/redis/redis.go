package redis

import (
	"context"
	"time"

	env "github.com/testground/learning-example/pkg/env"
	"github.com/testground/learning-example/pkg/util"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var client *redis.Client

func InitClient() {
	var err error
	for i := 10; i > 0; i-- {
		client = redis.NewClient(&redis.Options{
			Addr:     env.GetEnv().RedisUrl,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		pingCmd := client.Ping(ctx)
		_, err := pingCmd.Result()
		if err != nil {
			time.Sleep(time.Second * 2)
		} else {
			break
		}
	}

	// if err still exists, then we failed to connect to redis
	util.FailOnError(err, "Error connecting to Redis")
}

func GetClient() *redis.Client {
	if client == nil {
		InitClient()
	}
	return client
}
