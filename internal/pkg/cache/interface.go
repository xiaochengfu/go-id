package cache

import "time"

type Client interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
	Incr(key string) (interface{}, error)
	Close() error
}

//func NewClient() (Client,error){
//	isUserRedis := true
//	if isUserRedis {
//		return GetRedisInstance()
//	}else {
//		return nil,nil
//	}
//}
