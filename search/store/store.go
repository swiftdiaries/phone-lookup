package store

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// Pool is used to create a redis Pool
var Pool *redis.Pool

// NewPool is used to create a new redis pool
func NewPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Set is used to store a key, value pair in a Redis pool
func Set(key string, value string) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		panic(err)
	}
	return err
}

// Get is used to get a key's corresponding value
func Get(key string) string {
	conn := Pool.Get()
	defer conn.Close()

	resp, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return ""
	}
	return resp
}
