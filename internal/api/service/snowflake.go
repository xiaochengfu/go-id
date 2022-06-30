package service

type snowflake struct {
}

func (snow snowflake) GetId(business string) (interface{}, error) {
	return 8, nil
}
