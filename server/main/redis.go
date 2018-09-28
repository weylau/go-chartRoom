package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool


func initPool(addr string, maxIdles, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:maxIdles,
		MaxActive:maxActive,
		IdleTimeout:idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}