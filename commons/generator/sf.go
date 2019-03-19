package gs_commons_generator

import (
	"github.com/bwmarrin/snowflake"
	"zskparker.com/foundation/pkg/osenv"
)

func ID() *snowflake.Node {
	node, _ := snowflake.NewNode(osenv.GetNodeNumber())
	return node
}
