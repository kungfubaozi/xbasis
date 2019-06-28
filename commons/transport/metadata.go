package xbasistransport

import (
	"context"
	"github.com/micro/go-micro/metadata"
)

func NewContext(m map[string]string) context.Context {
	return metadata.NewContext(context.Background(), m)
}
