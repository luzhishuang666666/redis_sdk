package Redis_SDK

// SAdd 添加集合元素
func (db *RedisDB) SAdd(key string, val ...interface{}) error {
	err := db.client.SAdd(key, val).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// SCard 获取集合元素个数
func (db *RedisDB) SCard(key string) (int64, error) {
	size, err := db.client.SCard(key).Result()
	if err != nil {
		panic(err)
		return 0, err
	}
	return size, err
}

// SIsMember 判断元素是否在集合中
func (db *RedisDB) SIsMember(key string, val interface{}) (bool, error) {
	ok, err := db.client.SIsMember(key, val).Result()
	if err != nil {
		panic(err)
		return false, err
	}
	return ok, err
}

// SMembers 获取集合中所有的元素
func (db *RedisDB) SMembers(key string) ([]string, error) {
	es, err := db.client.SMembers(key).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return es, err
}

// SRem 删除集合元素
func (db *RedisDB) SRem(key string, members ...interface{}) error {
	err := db.client.SRem(key, members).Err()
	if err != nil {
		panic(err)
		return err
	}
	return err
}

// SPop 随机返回集合中的一个元素，并且删除这个元素
func (db *RedisDB) SPop(key string) (string, error) {
	val, err := db.client.SPop(key).Result()
	if err != nil {
		panic(err)
		return "", err
	}
	return val, err
}
