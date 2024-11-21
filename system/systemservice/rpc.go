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
