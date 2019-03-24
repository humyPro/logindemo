package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
)

var redisInit sync.Once
var conn redis.Conn

func initRedis() {
	fmt.Print("connecting to redis")
	_conn, e := redis.Dial("tcp", "127.0.0.1:6379")
	conn = _conn
	if e != nil {
		panic("Connect to redis error")
	}
}

func RedisCli() *redis.Conn {
	redisInit.Do(initRedis)
	return &conn
}
