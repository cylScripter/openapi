// Code generated by Kitex v0.11.3. DO NOT EDIT.

package systemservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	system "github.com/cylScripter/openapi/system"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, req *system.LoginReq, callOptions ...callopt.Option) (r *system.LoginResp, err error)
	CreateUser(ctx context.Context, req *system.CreateUserReq, callOptions ...callopt.Option) (r *system.CreateUserResp, err error)
	CreateRole(ctx context.Context, req *system.CreateRoleReq, callOptions ...callopt.Option) (r *system.CreateRoleResp, err error)
	CreatePermission(ctx context.Context, req *system.CreatePermissionReq, callOptions ...callopt.Option) (r *system.CreatePermissionResp, err error)
	SetUserRole(ctx context.Context, req *system.SetUserRoleReq, callOptions ...callopt.Option) (r *system.SetUserRoleResp, err error)
	SetRolePermission(ctx context.Context, req *system.SetRolePermissionReq, callOptions ...callopt.Option) (r *system.SetRolePermissionResp, err error)
	GetMenu(ctx context.Context, req *system.GetMenuReq, callOptions ...callopt.Option) (r *system.GetMenuResp, err error)
	GetUserInfo(ctx context.Context, req *system.GetUserInfoReq, callOptions ...callopt.Option) (r *system.GetUserInfoResp, err error)
	GetUserList(ctx context.Context, req *system.GetUserListReq, callOptions ...callopt.Option) (r *system.GetUserListResp, err error)
	GetAuthCode(ctx context.Context, req *system.GetAuthCodeReq, callOptions ...callopt.Option) (r *system.GetAuthCodeResp, err error)
	GetMenuList(ctx context.Context, req *system.GetMenuListReq, callOptions ...callopt.Option) (r *system.GetMenuListResp, err error)
	GetRoleList(ctx context.Context, req *system.GetRoleListReq, callOptions ...callopt.Option) (r *system.GetRoleListResp, err error)
	CreateMenu(ctx context.Context, req *system.CreateMenuReq, callOptions ...callopt.Option) (r *system.CreateMenuResp, err error)
	UpdateMenu(ctx context.Context, req *system.UpdateMenuReq, callOptions ...callopt.Option) (r *system.UpdateMenuResp, err error)
	GetApiList(ctx context.Context, req *system.GetApiListReq, callOptions ...callopt.Option) (r *system.GetApiListResp, err error)
	DeleteMenu(ctx context.Context, req *system.DeleteMenuReq, callOptions ...callopt.Option) (r *system.DeleteMenuResp, err error)
	GetAppList(ctx context.Context, req *system.GetAppListReq, callOptions ...callopt.Option) (r *system.GetAppListResp, err error)
	CreateApp(ctx context.Context, req *system.CreateAppReq, callOptions ...callopt.Option) (r *system.CreateAppResp, err error)
	GetApp(ctx context.Context, req *system.GetAppReq, callOptions ...callopt.Option) (r *system.GetAppResp, err error)
	CreateEduUser(ctx context.Context, req *system.CreateEduUserReq, callOptions ...callopt.Option) (r *system.CreateEduUserResp, err error)
	GetEduUserList(ctx context.Context, req *system.GetEduUserListReq, callOptions ...callopt.Option) (r *system.GetEduUserListResp, err error)
	GetPassword(ctx context.Context, req *system.GetPasswordReq, callOptions ...callopt.Option) (r *system.GetPasswordResp, err error)
	SetAppRoleMenu(ctx context.Context, req *system.SetAppRoleMenuReq, callOptions ...callopt.Option) (r *system.SetAppRoleMenuResp, err error)
	GetEduMenuList(ctx context.Context, req *system.GetEduMenuListReq, callOptions ...callopt.Option) (r *system.GetEduMenuListResp, err error)
	CreateEduMenu(ctx context.Context, req *system.CreateEduMenuReq, callOptions ...callopt.Option) (r *system.CreateEduMenuResp, err error)
	UpdateEduMenu(ctx context.Context, req *system.UpdateEduMenuReq, callOptions ...callopt.Option) (r *system.UpdateEduMenuResp, err error)
	DeleteEduMenu(ctx context.Context, req *system.DeleteEduMenuReq, callOptions ...callopt.Option) (r *system.DeleteEduMenuResp, err error)
	DeleteApp(ctx context.Context, req *system.DeleteAppReq, callOptions ...callopt.Option) (r *system.DeleteAppResp, err error)
	SetAppStatus(ctx context.Context, req *system.SetAppStatusReq, callOptions ...callopt.Option) (r *system.SetAppStatusResp, err error)
	SetEduMenuStatus(ctx context.Context, req *system.SetEduMenuStatusReq, callOptions ...callopt.Option) (r *system.SetEduMenuStatusResp, err error)
	GetOperationLogsList(ctx context.Context, req *system.GetOperationLogsListReq, callOptions ...callopt.Option) (r *system.GetOperationLogsListResp, err error)
	GetEduRoleList(ctx context.Context, req *system.GetEduRoleListReq, callOptions ...callopt.Option) (r *system.GetEduRoleListResp, err error)
	GetEduRoleMenu(ctx context.Context, req *system.GetEduRoleMenuReq, callOptions ...callopt.Option) (r *system.GetEduRoleMenuResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kSystemserviceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSystemserviceClient struct {
	*kClient
}

func (p *kSystemserviceClient) Login(ctx context.Context, req *system.LoginReq, callOptions ...callopt.Option) (r *system.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kSystemserviceClient) CreateUser(ctx context.Context, req *system.CreateUserReq, callOptions ...callopt.Option) (r *system.CreateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kSystemserviceClient) CreateRole(ctx context.Context, req *system.CreateRoleReq, callOptions ...callopt.Option) (r *system.CreateRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRole(ctx, req)
}

func (p *kSystemserviceClient) CreatePermission(ctx context.Context, req *system.CreatePermissionReq, callOptions ...callopt.Option) (r *system.CreatePermissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePermission(ctx, req)
}

func (p *kSystemserviceClient) SetUserRole(ctx context.Context, req *system.SetUserRoleReq, callOptions ...callopt.Option) (r *system.SetUserRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetUserRole(ctx, req)
}

func (p *kSystemserviceClient) SetRolePermission(ctx context.Context, req *system.SetRolePermissionReq, callOptions ...callopt.Option) (r *system.SetRolePermissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetRolePermission(ctx, req)
}

func (p *kSystemserviceClient) GetMenu(ctx context.Context, req *system.GetMenuReq, callOptions ...callopt.Option) (r *system.GetMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMenu(ctx, req)
}

func (p *kSystemserviceClient) GetUserInfo(ctx context.Context, req *system.GetUserInfoReq, callOptions ...callopt.Option) (r *system.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}

func (p *kSystemserviceClient) GetUserList(ctx context.Context, req *system.GetUserListReq, callOptions ...callopt.Option) (r *system.GetUserListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserList(ctx, req)
}

func (p *kSystemserviceClient) GetAuthCode(ctx context.Context, req *system.GetAuthCodeReq, callOptions ...callopt.Option) (r *system.GetAuthCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAuthCode(ctx, req)
}

func (p *kSystemserviceClient) GetMenuList(ctx context.Context, req *system.GetMenuListReq, callOptions ...callopt.Option) (r *system.GetMenuListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMenuList(ctx, req)
}

func (p *kSystemserviceClient) GetRoleList(ctx context.Context, req *system.GetRoleListReq, callOptions ...callopt.Option) (r *system.GetRoleListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRoleList(ctx, req)
}

func (p *kSystemserviceClient) CreateMenu(ctx context.Context, req *system.CreateMenuReq, callOptions ...callopt.Option) (r *system.CreateMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMenu(ctx, req)
}

func (p *kSystemserviceClient) UpdateMenu(ctx context.Context, req *system.UpdateMenuReq, callOptions ...callopt.Option) (r *system.UpdateMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMenu(ctx, req)
}

func (p *kSystemserviceClient) GetApiList(ctx context.Context, req *system.GetApiListReq, callOptions ...callopt.Option) (r *system.GetApiListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetApiList(ctx, req)
}

func (p *kSystemserviceClient) DeleteMenu(ctx context.Context, req *system.DeleteMenuReq, callOptions ...callopt.Option) (r *system.DeleteMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteMenu(ctx, req)
}

func (p *kSystemserviceClient) GetAppList(ctx context.Context, req *system.GetAppListReq, callOptions ...callopt.Option) (r *system.GetAppListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAppList(ctx, req)
}

func (p *kSystemserviceClient) CreateApp(ctx context.Context, req *system.CreateAppReq, callOptions ...callopt.Option) (r *system.CreateAppResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateApp(ctx, req)
}

func (p *kSystemserviceClient) GetApp(ctx context.Context, req *system.GetAppReq, callOptions ...callopt.Option) (r *system.GetAppResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetApp(ctx, req)
}

func (p *kSystemserviceClient) CreateEduUser(ctx context.Context, req *system.CreateEduUserReq, callOptions ...callopt.Option) (r *system.CreateEduUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEduUser(ctx, req)
}

func (p *kSystemserviceClient) GetEduUserList(ctx context.Context, req *system.GetEduUserListReq, callOptions ...callopt.Option) (r *system.GetEduUserListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEduUserList(ctx, req)
}

func (p *kSystemserviceClient) GetPassword(ctx context.Context, req *system.GetPasswordReq, callOptions ...callopt.Option) (r *system.GetPasswordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPassword(ctx, req)
}

func (p *kSystemserviceClient) SetAppRoleMenu(ctx context.Context, req *system.SetAppRoleMenuReq, callOptions ...callopt.Option) (r *system.SetAppRoleMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetAppRoleMenu(ctx, req)
}

func (p *kSystemserviceClient) GetEduMenuList(ctx context.Context, req *system.GetEduMenuListReq, callOptions ...callopt.Option) (r *system.GetEduMenuListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEduMenuList(ctx, req)
}

func (p *kSystemserviceClient) CreateEduMenu(ctx context.Context, req *system.CreateEduMenuReq, callOptions ...callopt.Option) (r *system.CreateEduMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEduMenu(ctx, req)
}

func (p *kSystemserviceClient) UpdateEduMenu(ctx context.Context, req *system.UpdateEduMenuReq, callOptions ...callopt.Option) (r *system.UpdateEduMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEduMenu(ctx, req)
}

func (p *kSystemserviceClient) DeleteEduMenu(ctx context.Context, req *system.DeleteEduMenuReq, callOptions ...callopt.Option) (r *system.DeleteEduMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteEduMenu(ctx, req)
}

func (p *kSystemserviceClient) DeleteApp(ctx context.Context, req *system.DeleteAppReq, callOptions ...callopt.Option) (r *system.DeleteAppResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteApp(ctx, req)
}

func (p *kSystemserviceClient) SetAppStatus(ctx context.Context, req *system.SetAppStatusReq, callOptions ...callopt.Option) (r *system.SetAppStatusResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetAppStatus(ctx, req)
}

func (p *kSystemserviceClient) SetEduMenuStatus(ctx context.Context, req *system.SetEduMenuStatusReq, callOptions ...callopt.Option) (r *system.SetEduMenuStatusResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetEduMenuStatus(ctx, req)
}

func (p *kSystemserviceClient) GetOperationLogsList(ctx context.Context, req *system.GetOperationLogsListReq, callOptions ...callopt.Option) (r *system.GetOperationLogsListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOperationLogsList(ctx, req)
}

func (p *kSystemserviceClient) GetEduRoleList(ctx context.Context, req *system.GetEduRoleListReq, callOptions ...callopt.Option) (r *system.GetEduRoleListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEduRoleList(ctx, req)
}

func (p *kSystemserviceClient) GetEduRoleMenu(ctx context.Context, req *system.GetEduRoleMenuReq, callOptions ...callopt.Option) (r *system.GetEduRoleMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEduRoleMenu(ctx, req)
}
