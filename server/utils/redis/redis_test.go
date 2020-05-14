package redis

import (
	"tasker/server/common"
	"testing"
)

func TestRedis(t *testing.T) {

	common.InitRedisConn()

	conn := common.GetRedisConn()

	_, _ = conn.Do("set", "hh", 1)
	defer conn.Close()

}
