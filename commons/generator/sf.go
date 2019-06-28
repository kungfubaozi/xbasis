package xbasisgenerator

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/twinj/uuid"
	"hash"
	"zskparker.com/foundation/pkg/osenv"
)

func ID() *snowflake.Node {
	node, _ := snowflake.NewNode(osenv.GetNodeNumber())
	return node
}

type IDGenerator interface {
	Get() string

	UUID() string

	String() string

	Short() string

	Node() *snowflake.Node
}

type generator struct {
	node *snowflake.Node
	h    hash.Hash
}

func (g *generator) UUID() string {
	g.h.Write([]byte(str(g.node.Generate().String())))
	return uuid.New([]byte(fmt.Sprintf("%x", g.h.Sum(nil)))).String()
}

func (g *generator) Short() string {
	return uuid.New([]byte(g.Get())).String()[24:]
}

func (g *generator) Node() *snowflake.Node {
	return g.node
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
