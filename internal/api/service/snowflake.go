package service

type snowflake struct {
}

func (snow snowflake) GetId(business string) interface{} {
	return 8
}
