package commonservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/common"
)

func UploadFile(ctx context.Context, req *common.UploadFileReq, callOptions ...client.Option) (resp *common.UploadFileResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.UploadFile(ctx, req)
}

func UploadNewMultipart(ctx context.Context, req *common.UploadNewMultipartReq, callOptions ...client.Option) (resp *common.UploadNewMultipartResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.UploadNewMultipart(ctx, req)
}

func GetPresignedUrlList(ctx context.Context, req *common.GetPresignedUrlListReq, callOptions ...client.Option) (resp *common.GetPresignedUrlListResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetPresignedUrlList(ctx, req)
}

func CompleteMultipart(ctx context.Context, req *common.CompleteMultipartReq, callOptions ...client.Option) (resp *common.CompleteMultipartResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.CompleteMultipart(ctx, req)
}

func AbortMultipart(ctx context.Context, req *common.AbortMultipartReq, callOptions ...client.Option) (resp *common.AbortMultipartResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.AbortMultipart(ctx, req)
}

func GetObject(ctx context.Context, req *common.GetObjectReq, callOptions ...client.Option) (resp *common.GetObjectResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetObject(ctx, req)
}

func DeleteObject(ctx context.Context, req *common.DeleteObjectReq, callOptions ...client.Option) (resp *common.DeleteObjectResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.DeleteObject(ctx, req)
}
