package Redis_SDK

import "github.com/go-redis/redis"

type RedisDB struct {
	client *redis.Client
	conf   RedisConf
}

func NewClient(conf RedisConf) *RedisDB {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
		PoolSize: conf.PoolSize,
	})

	if client == nil {
		return nil
	}

	return &RedisDB{
		client: client,
		conf:   conf,
	}
}
