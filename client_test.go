package Redis_SDK

import "testing"

var redisDB *RedisDB

func TestNewClient(t *testing.T) {
	conf := RedisConf{
		Addr:     "127.0.0.1",
		Password: "123456",
		Db:       1,
		PoolSize: 10,
	}

	redisDB = NewClient(conf)

	if redisDB == nil {
		t.Errorf("redis new client error")
		return
	}

	t.Log(redisDB.client.Ping().Result())

}
