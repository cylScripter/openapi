package helloservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/hello"
)

func SayHello(ctx context.Context, req *hello.SayHelloReq, callOptions ...client.Option) (resp *hello.SayHelloResp, err error) {
	clients := MustNewClient("helloservice", callOptions...)
	return clients.SayHello(ctx, req)
}
