package db

import (
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"github.com/xiaochengfu/go-id/configs"
	"log"
)

type cacheDriver struct {
	Client *redis.Client
}

func NewRedis() *cacheDriver {
	client, err := redisConnect()
	if err != nil {
		log.Fatal("redis连接异常")
	}
	return &cacheDriver{
		Client: client,
	}
}

func redisConnect() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass, // no password set
		DB:       cfg.Db,   // use default DB
	})
	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}
	return client, nil

}
