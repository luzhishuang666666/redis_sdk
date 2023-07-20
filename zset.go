package Redis_SDK

import "github.com/go-redis/redis"

// ZAdd 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (db *RedisDB) ZAdd(key string, z ...redis.Z) error {
	err := db.client.ZAdd(key, z...).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// ZCard 返回集合元素个数
func (db *RedisDB) ZCard(key string) (int64, error) {
	size, err := db.client.ZCard(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return size, err
}

// ZCount 统计某个分数范围内的元素个数
func (db *RedisDB) ZCount(key string, min string, max string) (int64, error) {
	size, err := db.client.ZCount(key, min, max).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return size, err
}

// ZIncrBy 增加元素的分数
func (db *RedisDB) ZIncrBy(key string, increment float64, member string) error {
	err := db.client.ZIncrBy(key, increment, member).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// ZRange 返回集合中某个索引范围的元素，根据分数从小到大排序
func (db *RedisDB) ZRange(key string, start, stop int64) ([]string, error) {
	vals, err := db.client.ZRange(key, start, stop).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return vals, err
}

// ZRangeByScore 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
func (db *RedisDB) ZRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	vals, err := db.client.ZRangeByScore(key, opt).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return vals, err
}

// ZRangeByScoreWithScores 除了返回集合元素，同时也返回元素对应的分数
func (db *RedisDB) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	vals, err := db.client.ZRangeByScoreWithScores(key, opt).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return vals, err
}

// ZRem 删除集合元素
func (db *RedisDB) ZRem(key string, vals ...string) error {
	err := db.client.ZRem(key, vals).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// ZRemRangeByRank 根据索引范围删除元素
func (db *RedisDB) ZRemRangeByRank(key string, start, stop int64) error {
	err := db.client.ZRemRangeByRank(key, start, stop).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// ZRemRangeByScore 根据分数范围删除元素
func (db *RedisDB) ZRemRangeByScore(key, min, max string) error {
	err := db.client.ZRemRangeByScore(key, min, max).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// ZScore 查询元素对应的分数
func (db *RedisDB) ZScore(key, member string) (float64, error) {
	score, err := db.client.ZScore(key, member).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return score, err
}

func (db *RedisDB) ZRank(key, member string) (int64, error) {
	rk, err := db.client.ZRank(key, member).Result()
	if err != nil {
		panic(err)
		return -1, err
	}
	return rk, err
}
