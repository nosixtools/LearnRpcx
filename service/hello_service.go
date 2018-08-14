package service

import (
	"context"
	"github.com/smallnest/rpcx/client"
	 rpcxplugins "github.com/nosixtools/rpcx-plugins/opentracing"

)

type HelloService struct {}

func (hs *HelloService) Hello(ctx context.Context, name *string , response *string) error {
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:7703", "")
	xclient := client.NewXClient("WorldService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	span, ctx, _ := rpcxplugins.GenerateSpanWithContext(ctx, "Hello")
	defer span.Finish()
	span.LogKV("step","service1")
	xclient.Call(ctx, "World", name, response)
	return nil
}
