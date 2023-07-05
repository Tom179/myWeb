package redis

import (
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func Connect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

/*
ctx := context.Background()

存err := rDB.Set(ctx, "key1", "hello", 1*time.Second).Err() //最后一个参数为储存过期时间,填0的意思是不设置过期时间（永久），而不是过期时间为0

取val, err := rDB.Get(ctx, "key1").Result()

*/
