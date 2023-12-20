package helper

import (
	"errors"
	"github.com/ugabiga/falcon/internal/ent/migrate"
	"strconv"
	"strings"
)

type NodeResolver struct {
}

func NewNodeResolver() *NodeResolver {
	return &NodeResolver{}
}

type DecodeNode struct {
	TableName string
	NodeID    int
}

func (r NodeResolver) Decode(id string) (DecodeNode, error) {
	splits := strings.Split(id, ":")
	if len(splits) != 2 {
		return DecodeNode{}, errors.New("invalid node id")
	}

	nodeTableName := splits[0]
	foundMatchedTable := false
	for _, table := range migrate.Tables {
		if table.Name == nodeTableName {
			foundMatchedTable = true
			break
		}
	}

	if !foundMatchedTable {
		return DecodeNode{}, errors.New("invalid node id")
	}

	rawNodeID := splits[1]
	parsedNodeID, err := strconv.Atoi(rawNodeID)
	if err != nil {
		return DecodeNode{}, err
	}

	return DecodeNode{
		TableName: nodeTableName,
		NodeID:    parsedNodeID,
	}, nil
}
