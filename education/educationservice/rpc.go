package educationservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/education"
)

func CreateApp(ctx context.Context, req *education.CreateAppReq, callOptions ...client.Option) (resp *education.CreateAppResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateApp(ctx, req)
}

func Login(ctx context.Context, req *education.LoginReq, callOptions ...client.Option) (resp *education.LoginResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.Login(ctx, req)
}

func GetUserList(ctx context.Context, req *education.GetUserListReq, callOptions ...client.Option) (resp *education.GetUserListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetUserList(ctx, req)
}

func GetRoleList(ctx context.Context, req *education.GetRoleListReq, callOptions ...client.Option) (resp *education.GetRoleListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetRoleList(ctx, req)
}

func GetPermissionList(ctx context.Context, req *education.GetPermissionListReq, callOptions ...client.Option) (resp *education.GetPermissionListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetPermissionList(ctx, req)
}

func GetMenuList(ctx context.Context, req *education.GetMenuReq, callOptions ...client.Option) (resp *education.GetMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetMenuList(ctx, req)
}
