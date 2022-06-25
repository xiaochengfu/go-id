package service

type IdSequence interface {
	GetId(businessKey string) interface{}
}

const (
	Seg  = "segment"
	Snow = "snowflake"
)

func NewIdSequence(model string) IdSequence {
	if model == Seg {
		return &segment{}
	} else if model == Snow {
		return &snowflake{}
	} else {
		panic("no know model")
	}
}
