package zredis

import (
    "errors"
    //"fmt"
    "github.com/garyburd/redigo/redis"
    "strings"
    "time"
)

type ZRedis struct {
    Servers   []string    // redis 服务器地址列表(包含port), 连接时依次尝试, 成功一次即可
    Sentinels []string    // 哨兵列表, 获取当前master
    Mtag      string      // 哨兵模式下的master标识
    Password  string      // 密码, 可为空
    DB        string      // DB编号
    AServer   string      // 当前使用的redis地址(A=active)
    Pool      *redis.Pool //连接池
}

func InitZRedis(servers, sentinels []string, passwd, db, mtag string) (zr *ZRedis, e error) {

    zr = &ZRedis{
        Servers:   servers,
        Sentinels: sentinels,
        Mtag:      mtag,
        Password:  passwd,
        DB:        db,
    }

    //获取可用的server
    if err := zr.GetActiveServer(); err != nil {
        return nil, err
    }

    //连接池
    zr.Pool = NewPool(zr.AServer, zr.Password, zr.DB)

    return
}

// 针对单个redis server建立连接池
func NewPool(server, password, db string) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     5,
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
}

//获取可用server
func (zr *ZRedis) GetActiveServer() error {
    if len(zr.Sentinels) > 0 { //有sentinel列表,向sentinel询问master
        for _, sentinel := range zr.Sentinels {
            //fmt.Println("sentinel:", sentinel)
            if c, err := redis.Dial("tcp", sentinel); err == nil {
                //连接成功, 找master
                mInfo, err := Strings(c.Do("SENTINEL", "get-master-addr-by-name", zr.Mtag))
                c.Close()
                //fmt.Println(mInfo)
                if err == nil && len(mInfo) == 2 { //host + port
                    zr.AServer = strings.Join(mInfo, ":")
                    return nil
                }
            }
        }
    } else {
        if len(zr.Servers) <= 0 {
            return errors.New("no server!")
        }
        for _, server := range zr.Servers {
            //尝试连接
            if c, err := redis.Dial("tcp", server); err == nil {
                //连接成功, 返回
                c.Close()
                zr.AServer = server
                return nil
            }
            //失败,继续
        }
    }
    return errors.New("Not Fond Server")
}
