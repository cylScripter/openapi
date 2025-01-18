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

func DeleteMenu(ctx context.Context, req *system.DeleteMenuReq, callOptions ...client.Option) (resp *system.DeleteMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.DeleteMenu(ctx, req)
}

func GetAppList(ctx context.Context, req *system.GetAppListReq, callOptions ...client.Option) (resp *system.GetAppListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetAppList(ctx, req)
}

func CreateApp(ctx context.Context, req *system.CreateAppReq, callOptions ...client.Option) (resp *system.CreateAppResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateApp(ctx, req)
}

func GetApp(ctx context.Context, req *system.GetAppReq, callOptions ...client.Option) (resp *system.GetAppResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetApp(ctx, req)
}

func CreateEduUser(ctx context.Context, req *system.CreateEduUserReq, callOptions ...client.Option) (resp *system.CreateEduUserResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateEduUser(ctx, req)
}

func GetEduUserList(ctx context.Context, req *system.GetEduUserListReq, callOptions ...client.Option) (resp *system.GetEduUserListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetEduUserList(ctx, req)
}

func GetPassword(ctx context.Context, req *system.GetPasswordReq, callOptions ...client.Option) (resp *system.GetPasswordResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetPassword(ctx, req)
}

func SetAppRoleMenu(ctx context.Context, req *system.SetAppRoleMenuReq, callOptions ...client.Option) (resp *system.SetAppRoleMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.SetAppRoleMenu(ctx, req)
}

func GetEduMenuList(ctx context.Context, req *system.GetEduMenuListReq, callOptions ...client.Option) (resp *system.GetEduMenuListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetEduMenuList(ctx, req)
}

func CreateEduMenu(ctx context.Context, req *system.CreateEduMenuReq, callOptions ...client.Option) (resp *system.CreateEduMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.CreateEduMenu(ctx, req)
}

func UpdateEduMenu(ctx context.Context, req *system.UpdateEduMenuReq, callOptions ...client.Option) (resp *system.UpdateEduMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.UpdateEduMenu(ctx, req)
}

func DeleteEduMenu(ctx context.Context, req *system.DeleteEduMenuReq, callOptions ...client.Option) (resp *system.DeleteEduMenuResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.DeleteEduMenu(ctx, req)
}

func DeleteApp(ctx context.Context, req *system.DeleteAppReq, callOptions ...client.Option) (resp *system.DeleteAppResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.DeleteApp(ctx, req)
}

func SetAppStatus(ctx context.Context, req *system.SetAppStatusReq, callOptions ...client.Option) (resp *system.SetAppStatusResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.SetAppStatus(ctx, req)
}

func SetEduMenuStatus(ctx context.Context, req *system.SetEduMenuStatusReq, callOptions ...client.Option) (resp *system.SetEduMenuStatusResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.SetEduMenuStatus(ctx, req)
}

func GetOperationLogsList(ctx context.Context, req *system.GetOperationLogsListReq, callOptions ...client.Option) (resp *system.GetOperationLogsListResp, err error) {
	clients := MustNewClient("system", callOptions...)
	return clients.GetOperationLogsList(ctx, req)
}
