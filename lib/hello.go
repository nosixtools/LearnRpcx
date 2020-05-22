package lib

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

type HelloService struct{}

func (hs *HelloService) Hello(ctx context.Context, name *string, response *string) error {

	span, ctx, _ := GenerateSpanWithContext(ctx, "Hello")
	defer span.Finish()
	fmt.Printf("%+v\n", *name)
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:7703", "")
	xclient := client.NewXClient("WorldService", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	span.LogKV("step", "service1")
	err := xclient.Call(ctx, "World", name, response)
	if err != nil {
		log.Fatal(err)
	}
	*response = "123"
	time.Sleep(10 * time.Millisecond)
	return nil
}
