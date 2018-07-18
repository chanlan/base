package midware

import (
	"testing"
	"github.com/garyburd/redigo/redis"
)

//TestConn
func TestConn(t *testing.T) {
	c := GetRedis()
	rs, _ := redis.String(c.Do("GET", "gome:sign:info:100040469104"))
	t.Log(rs)
	//c.Close()
}

//BenchmarkConn
func BenchmarkConn(b *testing.B) {
	b.N = 200
	c := GetRedis()
	if c.Err() != nil {
		b.Error(c.Err())
	}
}
