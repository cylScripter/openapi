package commonservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/common"
)

func GetFile(ctx context.Context, req *common.GetFileReq, callOptions ...client.Option) (resp *common.GetFileResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetFile(ctx, req)
}

func GetFileList(ctx context.Context, req *common.GetFileListReq, callOptions ...client.Option) (resp *common.GetFileListResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetFileList(ctx, req)
}

func UploadFile(ctx context.Context, req *common.UploadFileReq, callOptions ...client.Option) (resp *common.UploadFileResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.UploadFile(ctx, req)
}

func CompleteFile(ctx context.Context, req *common.CompleteFileReq, callOptions ...client.Option) (resp *common.CompleteFileResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.CompleteFile(ctx, req)
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

func CheckMultipartStatus(ctx context.Context, req *common.CheckMultipartStatusReq, callOptions ...client.Option) (resp *common.CheckMultipartStatusResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.CheckMultipartStatus(ctx, req)
}

func ResendPartUrl(ctx context.Context, req *common.ResendPartUrlReq, callOptions ...client.Option) (resp *common.ResendPartUrlResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.ResendPartUrl(ctx, req)
}

func GetUploadProgress(ctx context.Context, req *common.GetUploadProgressReq, callOptions ...client.Option) (resp *common.GetUploadProgressResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetUploadProgress(ctx, req)
}

func GetObject(ctx context.Context, req *common.GetObjectReq, callOptions ...client.Option) (resp *common.GetObjectResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetObject(ctx, req)
}

func DeleteObject(ctx context.Context, req *common.DeleteObjectReq, callOptions ...client.Option) (resp *common.DeleteObjectResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.DeleteObject(ctx, req)
}

func CreateAsyncTask(ctx context.Context, req *common.CreateAsyncTaskReq, callOptions ...client.Option) (resp *common.CreateAsyncTaskResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.CreateAsyncTask(ctx, req)
}

func GetAsyncTaskResults(ctx context.Context, req *common.GetAsyncTaskResultsReq, callOptions ...client.Option) (resp *common.GetAsyncTaskResultsResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.GetAsyncTaskResults(ctx, req)
}

func UpdateAsyncTask(ctx context.Context, req *common.UpdateAsyncTaskReq, callOptions ...client.Option) (resp *common.UpdateAsyncTaskResp, err error) {
	clients := MustNewClient("common", callOptions...)
	return clients.UpdateAsyncTask(ctx, req)
}
