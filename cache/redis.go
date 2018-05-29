package cache

import (
	"chat/conf"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitRedis() {
	redisConfig := conf.GetConfig().Redis
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisConfig.Address)
		},
	}
	return
}

func GetRedis() redis.Conn {
	return pool.Get()
}

func CloseRedis(conn redis.Conn) {
	conn.Close()
}
