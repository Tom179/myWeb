package redis

import (
	"fmt"
)

type RedisStore struct {
}

func (rs *RedisStore) Set(id string, value string) error {
	err := RDB.Set(ctx, id, value, 0).Err()
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return err
	}
	return nil
}

func (rs *RedisStore) Get(id string, clear bool) string {
	value, err := RDB.Get(ctx, id).Result()
	if err != nil {
		// 处理错误
	}
	if clear {
		err := RDB.Del(ctx, id).Err()
		if err != nil {
			fmt.Println(err)
			// 处理错误
		}
	}
	return value
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	storedAnswer := rs.Get(id, clear)
	return answer == storedAnswer
}
