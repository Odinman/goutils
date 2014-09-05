// 映射redigo的reply函数
package zredis

import (
    "github.com/garyburd/redigo/redis"
)

func Int(reply interface{}, err error) (int, error) {
    return redis.Int(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
    return redis.Int64(reply, err)
}

func Uint64(reply interface{}, err error) (uint64, error) {
    return redis.Uint64(reply, err)
}

func Float64(reply interface{}, err error) (float64, error) {
    return redis.Float64(reply, err)
}

func String(reply interface{}, err error) (string, error) {
    return redis.String(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
    return redis.Bytes(reply, err)
}

func Bool(reply interface{}, err error) (bool, error) {
    return redis.Bool(reply, err)
}

func Values(reply interface{}, err error) ([]interface{}, error) {
    return redis.Values(reply, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
    return redis.Strings(reply, err)
}
