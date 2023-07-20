package Redis_SDK

// LPush 从列表左边插入数据
func (db *RedisDB) LPush(key string, val ...interface{}) error {
	err := db.client.LPush(key, val...).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// LPushX 仅当列表存在的时候才插入数据
func (db *RedisDB) LPushX(key string, val interface{}) error {
	err := db.client.LPushX(key, val).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// RPop 从列表的右边删除第一个数据，并返回删除的数据
func (db *RedisDB) RPop(key string) (interface{}, error) {
	val, err := db.client.RPop(key).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return val, err

}

// RPush 从列表右边插入数据
func (db *RedisDB) RPush(key string, val ...interface{}) error {
	err := db.client.RPush(key, val).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// RPushX 仅当列表存在的时候才插入数据
func (db *RedisDB) RPushX(key string, val interface{}) error {
	err := db.client.RPushX(key, val).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func (db *RedisDB) LPop(key string) (interface{}, error) {
	val, err := db.client.LPop(key).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return val, err
}

// LLen 返回列表的大小
func (db *RedisDB) LLen(key string) (int64, error) {
	val, err := db.client.LLen(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return val, err
}

// LRem 删除列表中的数据
func (db *RedisDB) LRem(key string, count int64, value interface{}) (int64, error) {
	dels, err := db.client.LRem(key, count, value).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return dels, err
}

// LIndex 根据索引坐标，查询列表中的数据
func (db *RedisDB) LIndex(key string, index int64) (interface{}, error) {
	val, err := db.client.LIndex(key, index).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return val, err
}

// LInsert 在指定位置插入数据
func (db *RedisDB) LInsert(key, op string, pivot, value interface{}) error {
	err := db.client.LInsert(key, op, pivot, value).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}
