package main

import (
	"github.com/redis/go-redis/v9"
)

//type cacheDB struct {
//	redis redis.NewClient
//}

func redisConnection() {

}

func getConnection() {
	conf := getConfig()
	redis.NewClient(&redis.Options{
		Addr:     conf.Redis.host,
		Password: conf.Redis.password,
		DB:       0,
	})
}
