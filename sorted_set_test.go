package redisz

import (
	"testing"
)

func TestZadd(t *testing.T) {
	redisPool := NewRedisPool("hash", "192.168.1.141:6379", "", 1)

	testKey := "soretedset"

	testMap := make(map[int]string)
	testMap[1] = "USA"
	testMap[2] = "China"
	testMap[3] = "English"
	testMap[4] = "Japan"
	testMap[5] = "SSS"
	testMap[6] = "SSSA"
	ret := redisPool.Zadd(testKey, testMap)
	if ret < 0 {
		t.Error("Zadd failed.")
		return
	}

	score := redisPool.Zscore(testKey, "English")
	if score != 3 {
		t.Error("Zscore failed. score=", score)
		return
	}
	t.Log("return line=", ret, ", score=", score)

	if rets := redisPool.Zrange(testKey, 0, -1); rets != nil {
		t.Log("rets=", rets)
	}

	if rets := redisPool.ZrangeMap(testKey, 0, -1); rets != nil {
		t.Log("rets=", rets)
	}
}
