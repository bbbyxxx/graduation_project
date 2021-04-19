package dal

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	Agreement = "tcp"
	LocalHost = "localhost:6379"
	Pool      *redis.Pool
)

func init() {
	Pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         16,
		MaxActive:       0,
		IdleTimeout:     300,
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func GetConn() (redis.Conn, error) {
	c := Pool.Get()
	fmt.Println("redis conn success")
	return c, nil
}
