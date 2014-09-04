package zredis

import (
    "github.com/garyburd/redigo/redis"
    "time"
)

type Conn redis.Conn
type Pool redis.Pool

func NewPool(server, password, db string) *Pool {
    var p interface{}
    p = &redis.Pool{
        MaxIdle:     3,
        IdleTimeout: 240 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            if password != "" {
                if _, err := c.Do("AUTH", password); err != nil {
                    c.Close()
                    //beego.Debug("redis auth failed:", password)
                    return nil, err
                }
            }
            if db != "" {
                if _, err := c.Do("SELECT", db); err != nil {
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
    return p.(*Pool)
}
