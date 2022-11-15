package snowflakeId

import (
	"github.com/bwmarrin/snowflake"
)

type Twitter struct {
	node *snowflake.Node
}

func NewTwitter(workerId int64) (ISnowflakeId, error) {
	node, err := snowflake.NewNode(workerId)
	if err != nil {
		return nil, err
	}
	return &Twitter{node: node}, nil
}

func (sf *Twitter) Generate() ID {
	return sf.node.Generate()
}
