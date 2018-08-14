package service

import (
	"fmt"
	"context"
	rpcxplugins "github.com/nosixtools/rpcx-plugins/opentracing"

)

type WorldService struct {}

func (hs *WorldService) World(ctx context.Context, name *string , response *string) error {
	span, ctx, _ := rpcxplugins.GenerateSpanWithContext(ctx, "world")
	span.SetTag("method", "Hello")
	span.LogKV("step","rpcx")
	defer span.Finish()
	*response = fmt.Sprintf("name:%s,", *name)
	return nil
}
