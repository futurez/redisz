package redisz

import (
	"github.com/futurez/litego/logger"

	"github.com/garyburd/redigo/redis"
)

func (r *RedisPool) Zadd(key string, score int, member string) int {
	conn := r.pool.Get()
	defer conn.Close()

	val, err := redis.Int(conn.Do("ZADD", key, score, member))
	if err != nil {
		logger.Warn("ZADD ", r.server, " ", r.name, " ", err.Error())
		return -1
	}
	return val
}

func (r *RedisPool) Zadds(key string, scoremap map[int]string) int {
	conn := r.pool.Get()
	defer conn.Close()

	args := make([]interface{}, 0, len(scoremap)*2+1)
	args = append(args, key)
	for k, v := range scoremap {
		args = append(args, k, v)
	}

	val, err := redis.Int(conn.Do("ZADD", args...))
	if err != nil {
		logger.Warn("ZADD ", r.server, " ", r.name, " ", err.Error())
		return -1
	}
	return val
}

//get member score
func (r *RedisPool) Zscore(key string, member string) int {
	conn := r.pool.Get()
	defer conn.Close()

	val, err := redis.Int(conn.Do("ZSCORE", key, member))
	if err != nil {
		logger.Warn("ZADD ", r.server, " ", r.name, " ", err.Error())
		return -1
	}
	return val
}

//get member
func (r *RedisPool) Zrange(key string, start, end int) []string {
	conn := r.pool.Get()
	defer conn.Close()

	vals, err := redis.Strings(conn.Do("ZRANGE", key, start, end))
	if err != nil {
		logger.Warn("ZRANGE ", r.server, " ", r.name, " ", err.Error())
		return nil
	}
	return vals
}

func (r *RedisPool) ZrangeMap(key string, start, end int) map[string]string {
	conn := r.pool.Get()
	defer conn.Close()

	valMap, err := redis.StringMap(conn.Do("ZRANGE", key, start, end, "WITHSCORES"))
	if err != nil {
		logger.Warn("ZRANGE ... [WITHSCORES] ", r.server, " ", r.name, " ", err.Error())
		return nil
	}
	return valMap
}

//从大到小 get member
func (r *RedisPool) Zrevrrange(key string, start, end int) []string {
	conn := r.pool.Get()
	defer conn.Close()

	vals, err := redis.Strings(conn.Do("ZREVRRANGE", key, start, end))
	if err != nil {
		logger.Warn("ZREVRRANGE ", r.server, " ", r.name, " ", err.Error())
		return nil
	}
	return vals
}

func (r *RedisPool) ZrevrrangeMap(key string, start, end int) map[string]string {
	conn := r.pool.Get()
	defer conn.Close()

	valMap, err := redis.StringMap(conn.Do("ZREVRRANGE", key, start, end, "WITHSCORES"))
	if err != nil {
		logger.Warn("ZREVRRANGE ... [WITHSCORES] ", r.server, " ", r.name, " ", err.Error())
		return nil
	}
	return valMap
}
