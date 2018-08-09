/**
 * This is NOT a freeware, use is subject to license terms
 *
 * path   go-push/midware
 * date   2018/6/22 14:27 
 * author chenjingxiu
 */
package midware

import (
	"testing"
	"fmt"
	"reflect"
)

var predis *PRedis

func init() {
	predis = NewPRedis(RedisServers)
}

//TestRedisPool
func TestRedisPool(t *testing.T) {
	predis.GetRedis()
	predis.GetRedis()
	predis.GetRedis()
	fmt.Println("activeCount:", predis.Pool.ActiveCount())
	fmt.Println("idleCount:", predis.Pool.IdleCount())
	fmt.Println("Stats:", predis.Pool.Stats())
}

//TestSet
func TestSet(t *testing.T) {
	t.Log("------------SET---------------")
	rs, err := predis.Set("sign:info:100040469104", 1)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestGet
func TestGet(t *testing.T) {
	t.Log("------------GET---------------")
	rs, err := predis.Get("sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestDel
func TestDel(t *testing.T) {
	t.Log("------------DEL---------------")
	rs, err := predis.Del("sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestHSet
func TestHSet(t *testing.T) {
	t.Log("------------HSET---------------")
	rs, err := predis.HSet("new:sign:info:100040469104", "name2", "value2")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestHGet
func TestHGet(t *testing.T) {
	t.Log("------------HGET---------------")
	rs, err := predis.HGet("new:sign:info:100040469104", "name2")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestHGetALL
func TestHGetAll(t *testing.T) {
	t.Log("------------HGETALL---------------")
	rs, err := predis.HGetAll("new:sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestHMSet
func TestHMSet(t *testing.T) {
	t.Log("------------HMSET---------------")
	rs, err := predis.HMSet("new:sign:info:100040469104", "aa", "000", "bb", "1111")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestHDel
func TestHDel(t *testing.T) {
	t.Log("------------HDEL---------------")
	rs, err := predis.HDel("new:sign:info:100040469104", "name2")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestZAdd
func TestZAdd(t *testing.T) {
	t.Log("------------ZADD---------------")
	rs, err := predis.ZAdd("zadd:sign:info:100040469104", 0.2, "name4")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestZRange
func TestZRange(t *testing.T) {
	t.Log("------------ZRange---------------")
	rs, err := predis.ZRange("zadd:sign:info:100040469104", 0, -1)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestZCard
func TestZCard(t *testing.T) {
	t.Log("------------ZCard---------------")
	rs, err := predis.ZCard("zadd:sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestSAdd
func TestSAdd(t *testing.T) {
	t.Log("------------SAdd---------------")
	rs, err := predis.SAdd("sadd:sign:info:100040469104", "World")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestSMembers
func TestSMembers(t *testing.T) {
	t.Log("------------SMembers---------------")
	rs, err := predis.SMembers("sadd:sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestSRem
func TestSRem(t *testing.T) {
	t.Log("------------SRem---------------")
	rs, err := predis.SRem("sadd:sign:info:100040469104", "World")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestExpire
func TestExpire(t *testing.T) {
	t.Log("------------Expire---------------")
	predis.Set("exp:sign:info:100040469104", "1")
	rs, err := predis.Expire("exp:sign:info:100040469104", 60)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestTTL
func TestTTL(t *testing.T) {
	t.Log("------------TTL---------------")
	rs, err := predis.TTL("exp:sign:info:100040469104")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}

//TestINFO
func TestINFO(t *testing.T) {
	t.Log("------------INFO---------------")
	rs, err := predis.INFO()
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(reflect.TypeOf(rs))
	t.Log(rs)
}
