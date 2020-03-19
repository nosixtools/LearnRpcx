package main

import (
	"LearnRpcx/lib"
	"context"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
)

func main() {

	init := lib.Init("client")
	defer init.Close()

	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:7702", "")
	xclient := client.NewXClient("HelloService", client.Failfast, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	span, ctx, _ := lib.GenerateSpanWithContext(context.Background(), "start")
	defer span.Finish()

	// 设置标签
	span.SetTag("设置标签", "标签值")
	// 设置注释
	span.LogKV(
		"event", "soft error",
		"type", "cache timeout",
		"waited.millis", 1500)
	result := ""
	err := xclient.Call(ctx, "Hello", "nosix", &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
