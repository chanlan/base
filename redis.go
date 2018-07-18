package midware

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"strings"
	"errors"
)

var (
	redisPool *redis.Pool
	redisConn redis.Conn
)

const (
	RedisProtocol         = "tcp"
	RedisReadTimeout      = 1 * time.Minute
	RedisWriteTimeout     = 30 * time.Second
	RedisConnTimeout      = 1 * time.Minute
	RedisKeepAliveTimeout = 5 * time.Minute
	RedisDB               = 0
	RedisPwd              = ""
)

const (
	RedisMaxIdle         = 10
	RedisMaxActive       = 50
	RedisIdleTimeout     = 5 * time.Minute
	RedisWait            = false
	RedisMaxConnLifeTime = 1 * time.Hour
)

func init() {
	redisPool = NewRedisPool(RedisServers)
}

//NewRedisPool returns a new redis pool for the given redis server configuration.
func NewRedisPool(redisServers string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			for _, redisServer := range strings.Split(redisServers, ";") {
				conn, err := redis.Dial(RedisProtocol,
					redisServer,
					redis.DialReadTimeout(RedisReadTimeout),
					redis.DialConnectTimeout(RedisConnTimeout),
					redis.DialWriteTimeout(RedisWriteTimeout),
					redis.DialKeepAlive(RedisKeepAliveTimeout),
					redis.DialDatabase(RedisDB),
					redis.DialPassword(RedisPwd))
				if err == nil {
					return conn, nil
				} else {
					continue
				}
			}
			return nil, errors.New("redis conn error")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			return c.Err()
		},
		MaxIdle:         RedisMaxIdle,
		MaxActive:       RedisMaxActive,
		IdleTimeout:     RedisIdleTimeout,
		Wait:            RedisWait,
		MaxConnLifetime: RedisMaxConnLifeTime,
	}
}

//GetRedis returns a new redis connection to use from a redis pool
func GetRedis() redis.Conn {
	redisConn = redisPool.Get()
	return redisConn
}

//CloseRedis close a redis pool
func CloseRedisPool() error {
	return redisPool.Close()
}
