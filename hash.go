package Redis_SDK

// HSet 根据key和field字段设置，field字段的值
func (db *RedisDB) HSet(hashKey string, field string, fieldVal string) error {
	err := db.client.HSet(hashKey, field, fieldVal).Err()
	if err != nil {
		panic(err)
	}
	return err
}

// HGet 根据key和field字段，查询field字段的值
func (db *RedisDB) HGet(key string, field string) (string, error) {
	val, err := db.client.HGet(key, field).Result()
	if err != nil {
		panic(err)
		return "", err
	}
	return val, err
}

// HGetAll 一次性返回key=user_1的所有hash字段和值
func (db *RedisDB) HGetAll(key string) (map[string]string, error) {
	data, err := db.client.HGetAll(key).Result()
	if err != nil {
		panic(err)
		return nil, err
	}

	return data, err
}

// HIncrBy 根据key和field字段，累加字段的数值
func (db *RedisDB) HIncrBy(key string, field string) (int64, error) {
	count, err := db.client.HIncrBy(key, field, 2).Result()
	if err != nil {
		panic(err)
		return 0, err
	}

	return count, err
}

// HKeys 根据key返回所有字段名
func (db *RedisDB) HKeys(key string) ([]string, error) {
	keys, err := db.client.HKeys(key).Result()
	if err != nil {
		panic(err)
		return nil, err
	}

	return keys, err
}

// HLen 根据key，查询hash的字段数量
func (db *RedisDB) HLen(key string) (int64, error) {
	size, err := db.client.HLen(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return size, err
}

// HMGet 根据key和多个字段名，批量查询多个hash字段值
func (db *RedisDB) HMGet(key string, field ...string) ([]string, error) {
	vals, err := db.client.HMGet(key, field...).Result()
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

// HMSet 根据key和多个字段名和字段值，批量设置hash字段值
func (db *RedisDB) HMSet(key string, fields map[string]interface{}) error {
	err := db.client.HMSet(key, fields).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// HSetNX 如果field字段不存在，则设置hash字段值
func (db *RedisDB) HSetNX(key string, field string, val interface{}) error {
	err := db.client.HSetNX(key, field, val).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// HDel 根据key和字段名，删除hash字段，支持批量删除hash字段
func (db *RedisDB) HDel(key string, field ...string) error {
	err := db.client.HDel(key, field...).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

func (db *RedisDB) HExists(key string, field string) error {
	err := db.client.HExists(key, field).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}
