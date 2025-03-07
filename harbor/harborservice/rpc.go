package harborservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/harbor"
)

func Test(ctx context.Context, req *harbor.TestReq, callOptions ...client.Option) (resp *harbor.TestResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.Test(ctx, req)
}
