package app

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {

	//For connecting using unix socket
	//redisOpt := &redis.Options{
	//	Network:  "unix",
	//	Addr:     helper.Getenv("SOCKET", ""),
	//	Password: "",
	//	DB:       1,
	//}

	redisOpt := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	redisClient := redis.NewClient(redisOpt)
	return redisClient
}
