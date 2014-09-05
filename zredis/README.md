zredis
---

#功能

* 封装了redigo, https://github.com/garyburd/redigo
* 支持sentinel
* 支持pool

#用法

```
import (
    "github.com/zhaocloud/goutils/zredis"
)
var (
    ZhaoRedis     *zredis.ZRedis
)
func main() {
    ...
    ZhaoRedis, _ = zredis.InitZRedis(servers, sentinels, redisPwd, redisDB, redisMTag)
    redisConn := ZhaoRedis.Pool.Get()
    redisConn.Do(...)
    ...
}

```
