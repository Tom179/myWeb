package redis

import (
	"fmt"
)

type RedisStore struct { //用于实现Base64Captcha库的接口，相当于自己写了redis操作，不过封装在这个redisStore中
}

func (rs *RedisStore) Set(id string, value string) error { //过期时间为一分钟噢
	err := RDB.Set(ctx, id, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (rs *RedisStore) Get(id string, clear bool) string {
	value, err := RDB.Get(ctx, id).Result()
	if err != nil {
		fmt.Println("获取验证码失败")
		// 处理错误
	}
	if clear { //如果true，表示用完就删，即验证码为一次性的
		err := RDB.Del(ctx, id).Err()
		if err != nil {
			fmt.Println("验证码删除失败")
		}
	}
	return value
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	storedAnswer := rs.Get(id, clear)
	return answer == storedAnswer
}
