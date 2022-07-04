package service

type SegmentMemory struct{}

func (s SegmentMemory) GetId(businessKey string) (interface{}, error) {
	//todo 内存缓存待实现
	return nil, nil
}
