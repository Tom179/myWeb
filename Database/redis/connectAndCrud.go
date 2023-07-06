package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx context.Context

func Connect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx = context.Background()
}

func Create(key, value string) {
	if RDB == nil {
		Connect() //每次请求该接口都要连接和关闭，是否不合理？
	}
	if err := RDB.Set(ctx, key, value, 0).Err(); err != nil { //第三个参数设置过期时间
		fmt.Println(err)
		return
	}
}

func Delete(key string) {
	_, err := RDB.Del(ctx, key).Result()
	if err != nil {
		fmt.Println("删除键出错:", err)
		return
	}
}

/*
ctx := context.Background()

存err := rDB.Set(ctx, "key1", "hello", 1*time.Second).Err() //最后一个参数为储存过期时间,填0的意思是不设置过期时间（永久），而不是过期时间为0

取val, err := rDB.Get(ctx, "key1").Result()

*/
