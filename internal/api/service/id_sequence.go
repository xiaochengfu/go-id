package service

import "github.com/xiaochengfu/go-id/configs"

type Segment interface {
	GetId(businessKey string) (interface{}, error)
}

const (
	Seg          = "segment"
	Snow         = "snowflake"
	RedisDriver  = "redis"
	MemoryDriver = "memory"
)

func NewIdSequence() Segment {
	driver := configs.Get().Cache.Driver
	if driver == RedisDriver {
		return &SegmentRedis{}
	} else if driver == MemoryDriver {
		return &SegmentMemory{}
	} else {
		panic("no know model")
	}
}
