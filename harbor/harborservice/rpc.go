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

func GetProjectList(ctx context.Context, req *harbor.GetProjectListReq, callOptions ...client.Option) (resp *harbor.GetProjectListResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.GetProjectList(ctx, req)
}

func GetRepositoryList(ctx context.Context, req *harbor.GetRepositoryListReq, callOptions ...client.Option) (resp *harbor.GetRepositoryListResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.GetRepositoryList(ctx, req)
}

func GetArtifactList(ctx context.Context, req *harbor.GetArtifactListReq, callOptions ...client.Option) (resp *harbor.GetArtifactListResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.GetArtifactList(ctx, req)
}

func DeleteArtifact(ctx context.Context, req *harbor.DeleteArtifactReq, callOptions ...client.Option) (resp *harbor.DeleteArtifactResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.DeleteArtifact(ctx, req)
}

func GetModelHarborConfigList(ctx context.Context, req *harbor.GetHarborConfigListReq, callOptions ...client.Option) (resp *harbor.GetHarborConfigListResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.GetModelHarborConfigList(ctx, req)
}

func DeleteHarborConfig(ctx context.Context, req *harbor.DeleteHarborConfigReq, callOptions ...client.Option) (resp *harbor.DeleteHarborConfigResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.DeleteHarborConfig(ctx, req)
}

func CreateHarborConfig(ctx context.Context, req *harbor.CreateHarborConfigReq, callOptions ...client.Option) (resp *harbor.CreateHarborConfigResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.CreateHarborConfig(ctx, req)
}

func UpdateHarborConfig(ctx context.Context, req *harbor.UpdateHarborConfigReq, callOptions ...client.Option) (resp *harbor.UpdateHarborConfigResp, err error) {
	clients := MustNewClient("harbor", callOptions...)
	return clients.UpdateHarborConfig(ctx, req)
}
