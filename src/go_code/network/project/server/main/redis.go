package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

// "119.45.252.206:6379"
func initPool(address string, maxIdle, maxActive int, idleTime time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTime,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
