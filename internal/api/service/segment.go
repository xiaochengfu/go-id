package service

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	"github.com/xiaochengfu/go-id/internal/pkg/db"
)

type segment struct{}

func (s segment) GetId(businessKey string) (interface{}, error) {
	client := db.NewRedis().Client
	defer client.Close()

	r := client.Get(businessKey)
	switch {
	case r.Err() == redis.Nil, r.Val() == "":
		var initNum int64
		initNum = 1
		fmt.Println("value is empty || not exist")
		status := client.Set(businessKey, initNum, -1)
		fmt.Println(status)
		return initNum, nil
	case r.Err() != nil:
		fmt.Println("Get failed", r.Err())
	default:
		fmt.Println("value incr")
		num := client.Incr(businessKey)
		return num.Val(), nil
	}
	fmt.Println("end")
	err := errors.New("get id error")
	return nil, err
}
