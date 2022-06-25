package service

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type segment struct{}

func (s segment) GetId(businessKey string) interface{} {
	c, err := redis.Dial("tcp", "59.110.213.203:6379", redis.DialDatabase(1), redis.DialPassword("hp8805623"))
	if err != nil {
		fmt.Println("conn redis failed,", err)
	}

	defer c.Close()
	r, err := c.Do("Get", businessKey)
	if err != nil {
		fmt.Println("get abc failed,", err)
	}
	if r != nil {
		r, err = c.Do("Incr", businessKey)
		if err != nil {
			fmt.Println(r)
		}
	} else {
		var initNum int64
		initNum = 1
		_, err = c.Do("Set", businessKey, initNum)
		if err != nil {
			fmt.Println(err)
		}
		return initNum
	}
	return r.(int64)
}
