package Redis_SDK

import (
	"fmt"
	"time"
)

// Set Redis的set命令，不含有expiration
func (db *RedisDB) Set(key string, value string) error {
	err := db.client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// SetByExpiration Redis的set命令，含有expiration
func (db *RedisDB) SetByExpiration(key string, value string, expiration time.Duration) error {
	err := db.client.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// Get Redis的get命令，返回value和err
func (db *RedisDB) Get(key string) (string, error) {
	val, err := db.client.Get(key).Result()

	if err != nil {
		panic(err)
		return "", err
	}

	return val, err
}

// GetSet 设置一个key的值，并返回这个key的旧值
func (db *RedisDB) GetSet(key string, newValue string) (string, error) {
	oldVal, err := db.client.GetSet(key, newValue).Result()

	if err != nil {
		panic(err)
		return "", err
	}
	return oldVal, err
}

// SetNX 如果key不存在，则设置这个key的值
func (db *RedisDB) SetNX(key string, value string) error {
	err := db.client.SetNX(key, value, 0).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// SetNXByExpiration 如果key不存在，则设置这个key的值,增加expiration
func (db *RedisDB) SetNXByExpiration(key string, value string, expiration time.Duration) error {
	err := db.client.SetNX(key, value, expiration).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// MGet 批量查询key的值
func (db *RedisDB) MGet(keys ...string) ([]string, error) {
	vals, err := db.client.MGet(keys...).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	strVals := make([]string, len(vals))
	for i, v := range vals {
		if v == nil {
			strVals[i] = ""
		} else {
			strVals[i] = v.(string)
		}
	}
	return strVals, err
}

// MSet 批量设置key的值
func (db *RedisDB) MSet(keyValues ...string) error {
	if len(keyValues)%2 != 0 {
		return fmt.Errorf("MSet: number of args must be even")
	}
	kvs := make(map[string]interface{})
	for i := 0; i < len(keyValues); i += 2 {
		kvs[keyValues[i]] = keyValues[i+1]
	}
	_, err := db.client.MSet(kvs).Result()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// Incr Incr函数每次加一
func (db *RedisDB) Incr(key string) (int64, error) {
	val, err := db.client.Incr(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return val, err
}

// IncrBy IncrBy函数，可以指定每次递增多少
func (db *RedisDB) IncrBy(key string, by int64) (int64, error) {
	val, err := db.client.IncrBy(key, by).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return val, err
}

// Decr Decr函数每次减一
func (db *RedisDB) Decr(key string) (int64, error) {
	val, err := db.client.Decr(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}

	return val, err
}

// DecByr DecrBy函数，可以指定每次递减多少
func (db *RedisDB) DecByr(key string, by int64) (int64, error) {
	val, err := db.client.DecrBy(key, by).Result()
	if err != nil {
		panic(err)
		return 0, err
	}

	return val, err
}

// Del 删除多个key, Del函数支持删除多个key
func (db *RedisDB) Del(key ...string) error {
	err := db.client.Del(key...).Err()
	if err != nil {
		panic(err)
		return err
	}

	return err
}

// Expire 设置过期时间
func (db *RedisDB) Expire(key string, sec int) error {
	err := db.Expire(key, sec)
	if err != nil {
		panic(err)
		return err
	}

	return err
}
