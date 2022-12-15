package main

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func httpServer(conn redis.Conn, pool *redis.Pool, mode int) {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		stats := pool.Stats()
		str := fmt.Sprintf("%+v", stats)
		w.Write([]byte(str))
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		var err error
		var rs interface{}
		if mode == 1 {
			rs, err = conn.Do("ping")
		} else {
			ctx, _ := context.WithTimeout(context.Background(), time.Minute)
			rs, err = redis.DoContext(conn, ctx, "ping")
		}
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(rs.(string)))
		}
	})
	http.ListenAndServe("0.0.0.0:6060", nil)
}

func getPoll() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     1,
		MaxActive:   3,
		IdleTimeout: 3600 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", 1); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func testDo() {
	redisPool := getPoll()
	c := redisPool.Get()
	httpServer(c, redisPool, 1)
}

func testDoContext() {
	redisPool := getPoll()
	c := redisPool.Get()
	httpServer(c, redisPool, 2)
}

func main() {
	// conn.Do是直接设置net.conn的超时
	testDo()

	// redis.DoContext是主协程select，新协程执行conn.Read
	testDoContext()
}
