/*
 * @Author: wtf
 * @Date: 2020-08-28 19:10:31
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 20:24:57
 * @Description: plase write Description
 */
package gredis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/tengfei31/website/pkg/setting"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle: setting.RedisSetting.MaxIdle,
		MaxActive: setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				_, err = c.Do("AUTH", setting.RedisSetting.Password)
				if err != nil {
					defer c.Close()
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func (c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DELETE", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("*%s*", key)))
	if err != nil {
		return err
	}
	for _, tmpKey := range keys {
		_, err = Delete(tmpKey)
		if err != nil {
			return err
		}
	}
	return nil
}







