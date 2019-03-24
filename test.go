package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func main() {
	_conn, e := redis.Dial("tcp", "127.0.0.1:6379")
	conn = _conn
	if e != nil {
		fmt.Println("Connect to redis error", e)
		return
	}

	reply, e := conn.Do("DEL", "xxx")
	println(reply)
	println(e)
}

type A struct {
	a int
	b string
}
