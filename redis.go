/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/22 10:40
 * author chenjingxiu
 */
package midware

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"strings"
	"errors"
	"sync"
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
	RedisMaxActive       = 5
	RedisIdleTimeout     = 240 * time.Second
	RedisWait            = false
	RedisMaxConnLifeTime = 1 * time.Hour
)

var lock sync.Mutex

//newRedisPool returns a new redis pool for the given redis server configuration.
func newRedisPool(redisServers string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			for _, redisServer := range strings.Split(redisServers, ";") {
				conn, err := redis.Dial(RedisProtocol, redisServer,
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
			return nil, errors.New("redis connection error, server info:[" + redisServers + "]")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < RedisConnTimeout {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         RedisMaxIdle,
		MaxActive:       RedisMaxActive,
		IdleTimeout:     RedisIdleTimeout,
		Wait:            RedisWait,
		MaxConnLifetime: RedisMaxConnLifeTime,
	}
}

type PRedis struct {
	Prefix  string
	Postfix string
	Pool    *redis.Pool
}

func NewPRedis(redisServers string, options ...string) *PRedis {
	pool := newRedisPool(redisServers)
	if len(options) > 1 {
		return &PRedis{
			Prefix:  options[0],
			Postfix: options[1],
			Pool:    pool,
		}
	} else if len(options) > 0 {
		return &PRedis{
			Prefix: options[0],
			Pool:   pool,
		}
	}
	return &PRedis{
		Pool: pool,
	}
}

//GetRedis returns a new redis connection to use from a redis pool
func (this *PRedis) GetRedis() redis.Conn {
	return this.Pool.Get()
}

func (this *PRedis) appendPrefix(prefix string) {
	this.Prefix = prefix
}

func (this *PRedis) appendPostfix(postfix string) {
	this.Postfix = postfix
}

func (this *PRedis) getKey(oldKey string) string {
	return this.Prefix + oldKey + this.Postfix
}

func (this *PRedis) Get(oldKey string) (string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return rs, nil
}

func (this *PRedis) Set(oldKey string, data interface{}) (string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.String(conn.Do("SET", key, data))
	if err != nil {
		return "", err
	}
	return rs, nil
}

func (this *PRedis) Del(oldKey string) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("DEL", key))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) HSet(oldKey string, name interface{}, value interface{}) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("HSET", key, name, value))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) HMSet(oldKey string, data ...interface{}) (string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	lock.Lock()
	var param []interface{}
	param = append(param, key)
	for k, v := range data {
		param = append(param, k, v)
	}
	rs, err := redis.String(conn.Do("HMSET", param...))
	lock.Unlock()
	if err != nil {
		return "", err
	}
	return rs, nil
}

func (this *PRedis) HDel(oldKey string, name string) (interface{}, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := conn.Do("HDEL", key, name)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (this *PRedis) HGet(oldKey string, name interface{}) (string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.String(conn.Do("HGET", key, name))
	if err != nil {
		return "", err
	}
	return rs, nil
}

func (this *PRedis) HGetAll(oldKey string) (map[string]string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (this *PRedis) ZAdd(oldKey string, score interface{}, data interface{}) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("ZADD", key, score, data))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) ZCard(oldKey string) (int, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int(conn.Do("ZCARD", key))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) ZRange(oldKey string, start interface{}, end interface{}) ([]string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Strings(conn.Do("ZRANGE", key, start, end))
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (this *PRedis) SAdd(oldKey string, data interface{}) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("SADD", key, data))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) SRem(oldKey string, data interface{}) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("SREM", key, data))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) SMembers(oldKey string) ([]string, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (this *PRedis) Expire(oldKey string, expireAt int64) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("EXPIRE", key, expireAt))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) TTL(oldKey string) (int64, error) {
	key := this.getKey(oldKey)
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.Int64(conn.Do("TTL", key))
	if err != nil {
		return 0, err
	}
	return rs, nil
}

func (this *PRedis) INFO() (string, error) {
	conn := this.GetRedis()
	defer conn.Close()
	rs, err := redis.String(conn.Do("INFO"))
	if err != nil {
		return "", err
	}
	return rs, nil
}

//Close redis pool, release resource
//Don't use it if you are not need to do
func (this *PRedis) Close() error {
	return this.Pool.Close()
}
