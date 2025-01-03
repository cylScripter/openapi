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

func CreateUser(ctx context.Context, req *education.CreateUserReq, callOptions ...client.Option) (resp *education.CreateUserResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateUser(ctx, req)
}

func CreateRole(ctx context.Context, req *education.CreateRoleReq, callOptions ...client.Option) (resp *education.CreateRoleResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateRole(ctx, req)
}

func DeleteUser(ctx context.Context, req *education.DeleteUserReq, callOptions ...client.Option) (resp *education.DeleteUserResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteUser(ctx, req)
}

func GetUserList(ctx context.Context, req *education.GetUserListReq, callOptions ...client.Option) (resp *education.GetUserListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetUserList(ctx, req)
}

func GetRoleList(ctx context.Context, req *education.GetRoleListReq, callOptions ...client.Option) (resp *education.GetRoleListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetRoleList(ctx, req)
}

func CreatePermission(ctx context.Context, req *education.CreatePermissionReq, callOptions ...client.Option) (resp *education.CreatePermissionResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreatePermission(ctx, req)
}

func GetPermissionList(ctx context.Context, req *education.GetPermissionListReq, callOptions ...client.Option) (resp *education.GetPermissionListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetPermissionList(ctx, req)
}

func GetMenu(ctx context.Context, req *education.GetMenuReq, callOptions ...client.Option) (resp *education.GetMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetMenu(ctx, req)
}

func GetMenuList(ctx context.Context, req *education.GetMenuReq, callOptions ...client.Option) (resp *education.GetMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetMenuList(ctx, req)
}

func SetUserRole(ctx context.Context, req *education.SetUserRoleReq, callOptions ...client.Option) (resp *education.SetUserRoleResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SetUserRole(ctx, req)
}

func SetRolePermission(ctx context.Context, req *education.SetRolePermissionReq, callOptions ...client.Option) (resp *education.SetRolePermissionResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SetRolePermission(ctx, req)
}

func SetRoleMenu(ctx context.Context, req *education.SetRoleMenuReq, callOptions ...client.Option) (resp *education.SetRoleMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SetRoleMenu(ctx, req)
}

func CreateMenu(ctx context.Context, req *education.CreateMenuReq, callOptions ...client.Option) (resp *education.CreateMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateMenu(ctx, req)
}

func CreateOffice(ctx context.Context, req *education.CreateOfficeReq, callOptions ...client.Option) (resp *education.CreateOfficeResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateOffice(ctx, req)
}

func GetOfficeList(ctx context.Context, req *education.GetOfficeListReq, callOptions ...client.Option) (resp *education.GetOfficeListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetOfficeList(ctx, req)
}