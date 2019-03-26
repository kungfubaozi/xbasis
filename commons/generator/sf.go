package gs_commons_generator

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"hash"
	"zskparker.com/foundation/pkg/osenv"
)

func ID() *snowflake.Node {
	node, _ := snowflake.NewNode(osenv.GetNodeNumber())
	return node
}

type IDGenerator interface {
	Get() string

	String() string
}

type generator struct {
	node *snowflake.Node
	h    hash.Hash
}

func (g *generator) String() string {
	return g.node.Generate().String()
}

func str(str string) string {
	var b string
	for _, c := range str {
		b = fmt.Sprint(b, c)
	}
	return b
}

func (g *generator) Get() string {
	g.h.Write([]byte(str(g.node.Generate().String())))
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", g.h.Sum(nil))))
}

func NewIDG() IDGenerator {
	h := sha1.New()
	return &generator{node: ID(), h: h}
}
