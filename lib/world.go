package lib

import (
	"context"
	"fmt"
)

type WorldService struct{}

func (hs *WorldService) World(ctx context.Context, name *string, response *string) error {
	span, ctx, _ := GenerateSpanWithContext(ctx, "world")
	defer span.Finish()

	span.SetTag("method", "Hello")
	span.LogKV("step", "rpcx")

	*response = fmt.Sprintf("name:%s,", *name)
	return nil
}
