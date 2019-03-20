package gs_commons_transport

import (
	"context"
	"github.com/micro/go-micro/metadata"
)

func Verify() error {

}

func InsideContext(service string) context.Context {
	return metadata.NewContext(context.Background(), map[string]string{"X-Micro-Gosion": service})
}
