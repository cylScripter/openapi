package systemservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/system"
)

func Login(ctx context.Context, req *system.LoginReq, callOptions ...client.Option) (resp *system.LoginResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.Login(ctx, req)
}
