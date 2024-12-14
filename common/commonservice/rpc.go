package commonservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/common"
)

func UploadImage(ctx context.Context, req *common.UploadImageReq, callOptions ...client.Option) (resp *common.UploadImageResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.UploadImage(ctx, req)
}
