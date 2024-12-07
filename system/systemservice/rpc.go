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

func CreateUser(ctx context.Context, req *system.CreateUserReq, callOptions ...client.Option) (resp *system.CreateUserResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateUser(ctx, req)
}

func CreateRole(ctx context.Context, req *system.CreateRoleReq, callOptions ...client.Option) (resp *system.CreateRoleResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateRole(ctx, req)
}

func CreatePermission(ctx context.Context, req *system.CreatePermissionReq, callOptions ...client.Option) (resp *system.CreatePermissionResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreatePermission(ctx, req)
}

func SetUserRole(ctx context.Context, req *system.SetUserRoleReq, callOptions ...client.Option) (resp *system.SetUserRoleResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.SetUserRole(ctx, req)
}

func SetRolePermission(ctx context.Context, req *system.SetRolePermissionReq, callOptions ...client.Option) (resp *system.SetRolePermissionResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.SetRolePermission(ctx, req)
}

func GetMenu(ctx context.Context, req *system.GetMenuReq, callOptions ...client.Option) (resp *system.GetMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetMenu(ctx, req)
}

func GetUserInfo(ctx context.Context, req *system.GetUserInfoReq, callOptions ...client.Option) (resp *system.GetUserInfoResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetUserInfo(ctx, req)
}

func GetUserList(ctx context.Context, req *system.GetUserListReq, callOptions ...client.Option) (resp *system.GetUserListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetUserList(ctx, req)
}

func GetAuthCode(ctx context.Context, req *system.GetAuthCodeReq, callOptions ...client.Option) (resp *system.GetAuthCodeResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetAuthCode(ctx, req)
}

func GetMenuList(ctx context.Context, req *system.GetMenuListReq, callOptions ...client.Option) (resp *system.GetMenuListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetMenuList(ctx, req)
}

func GetRoleList(ctx context.Context, req *system.GetRoleListReq, callOptions ...client.Option) (resp *system.GetRoleListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetRoleList(ctx, req)
}

func CreateMenu(ctx context.Context, req *system.CreateMenuReq, callOptions ...client.Option) (resp *system.CreateMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateMenu(ctx, req)
}

func UpdateMenu(ctx context.Context, req *system.UpdateMenuReq, callOptions ...client.Option) (resp *system.UpdateMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.UpdateMenu(ctx, req)
}

func GetApiList(ctx context.Context, req *system.GetApiListReq, callOptions ...client.Option) (resp *system.GetApiListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetApiList(ctx, req)
}
