package common

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var conn redis.Conn

func InitRedisConn() redis.Conn {
	address := viper.GetString("redis.address")
	network := viper.GetString("redis.network")
	c, err := redis.Dial(network, address)
	if err != nil {
		panic(err)
	}
	conn = c
	return c
}

func GetRedisConn() redis.Conn {
	return conn
}
