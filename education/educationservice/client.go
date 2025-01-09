// Code generated by Kitex v0.11.3. DO NOT EDIT.

package educationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	education "github.com/cylScripter/openapi/education"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateApp(ctx context.Context, req *education.CreateAppReq, callOptions ...callopt.Option) (r *education.CreateAppResp, err error)
	Login(ctx context.Context, req *education.LoginReq, callOptions ...callopt.Option) (r *education.LoginResp, err error)
	CreateUser(ctx context.Context, req *education.CreateUserReq, callOptions ...callopt.Option) (r *education.CreateUserResp, err error)
	CreateRole(ctx context.Context, req *education.CreateRoleReq, callOptions ...callopt.Option) (r *education.CreateRoleResp, err error)
	DeleteUser(ctx context.Context, req *education.DeleteUserReq, callOptions ...callopt.Option) (r *education.DeleteUserResp, err error)
	GetUserList(ctx context.Context, req *education.GetUserListReq, callOptions ...callopt.Option) (r *education.GetUserListResp, err error)
	GetRoleList(ctx context.Context, req *education.GetRoleListReq, callOptions ...callopt.Option) (r *education.GetRoleListResp, err error)
	CreatePermission(ctx context.Context, req *education.CreatePermissionReq, callOptions ...callopt.Option) (r *education.CreatePermissionResp, err error)
	GetPermissionList(ctx context.Context, req *education.GetPermissionListReq, callOptions ...callopt.Option) (r *education.GetPermissionListResp, err error)
	GetMenu(ctx context.Context, req *education.GetMenuReq, callOptions ...callopt.Option) (r *education.GetMenuResp, err error)
	GetMenuList(ctx context.Context, req *education.GetMenuListReq, callOptions ...callopt.Option) (r *education.GetMenuListResp, err error)
	SetUserRole(ctx context.Context, req *education.SetUserRoleReq, callOptions ...callopt.Option) (r *education.SetUserRoleResp, err error)
	SetRolePermission(ctx context.Context, req *education.SetRolePermissionReq, callOptions ...callopt.Option) (r *education.SetRolePermissionResp, err error)
	SetRoleMenu(ctx context.Context, req *education.SetRoleMenuReq, callOptions ...callopt.Option) (r *education.SetRoleMenuResp, err error)
	CreateMenu(ctx context.Context, req *education.CreateMenuReq, callOptions ...callopt.Option) (r *education.CreateMenuResp, err error)
	CreateOffice(ctx context.Context, req *education.CreateOfficeReq, callOptions ...callopt.Option) (r *education.CreateOfficeResp, err error)
	GetOfficeList(ctx context.Context, req *education.GetOfficeListReq, callOptions ...callopt.Option) (r *education.GetOfficeListResp, err error)
	ImportCourseApply(ctx context.Context, req *education.ImportCourseApplyReq, callOptions ...callopt.Option) (r *education.ImportCourseApplyResp, err error)
	GetCourseApplyList(ctx context.Context, req *education.GetCourseApplyListReq, callOptions ...callopt.Option) (r *education.GetCourseApplyListResp, err error)
	GetSelfCourseApplyList(ctx context.Context, req *education.GetSelfCourseApplyListReq, callOptions ...callopt.Option) (r *education.GetSelfCourseApplyListResp, err error)
	DeleteCourseApply(ctx context.Context, req *education.DeleteCourseApplyReq, callOptions ...callopt.Option) (r *education.DeleteCourseApplyResp, err error)
	ApproveCourseApply(ctx context.Context, req *education.ApproveCourseApplyReq, callOptions ...callopt.Option) (r *education.ApproveCourseApplyResp, err error)
	ResetCourseApply(ctx context.Context, req *education.ResetCourseApplyReq, callOptions ...callopt.Option) (r *education.ResetCourseApplyResp, err error)
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
	return &kEducationserviceClient{
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

type kEducationserviceClient struct {
	*kClient
}

func (p *kEducationserviceClient) CreateApp(ctx context.Context, req *education.CreateAppReq, callOptions ...callopt.Option) (r *education.CreateAppResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateApp(ctx, req)
}

func (p *kEducationserviceClient) Login(ctx context.Context, req *education.LoginReq, callOptions ...callopt.Option) (r *education.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kEducationserviceClient) CreateUser(ctx context.Context, req *education.CreateUserReq, callOptions ...callopt.Option) (r *education.CreateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kEducationserviceClient) CreateRole(ctx context.Context, req *education.CreateRoleReq, callOptions ...callopt.Option) (r *education.CreateRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRole(ctx, req)
}

func (p *kEducationserviceClient) DeleteUser(ctx context.Context, req *education.DeleteUserReq, callOptions ...callopt.Option) (r *education.DeleteUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, req)
}

func (p *kEducationserviceClient) GetUserList(ctx context.Context, req *education.GetUserListReq, callOptions ...callopt.Option) (r *education.GetUserListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserList(ctx, req)
}

func (p *kEducationserviceClient) GetRoleList(ctx context.Context, req *education.GetRoleListReq, callOptions ...callopt.Option) (r *education.GetRoleListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRoleList(ctx, req)
}

func (p *kEducationserviceClient) CreatePermission(ctx context.Context, req *education.CreatePermissionReq, callOptions ...callopt.Option) (r *education.CreatePermissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePermission(ctx, req)
}

func (p *kEducationserviceClient) GetPermissionList(ctx context.Context, req *education.GetPermissionListReq, callOptions ...callopt.Option) (r *education.GetPermissionListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPermissionList(ctx, req)
}

func (p *kEducationserviceClient) GetMenu(ctx context.Context, req *education.GetMenuReq, callOptions ...callopt.Option) (r *education.GetMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMenu(ctx, req)
}

func (p *kEducationserviceClient) GetMenuList(ctx context.Context, req *education.GetMenuListReq, callOptions ...callopt.Option) (r *education.GetMenuListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMenuList(ctx, req)
}

func (p *kEducationserviceClient) SetUserRole(ctx context.Context, req *education.SetUserRoleReq, callOptions ...callopt.Option) (r *education.SetUserRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetUserRole(ctx, req)
}

func (p *kEducationserviceClient) SetRolePermission(ctx context.Context, req *education.SetRolePermissionReq, callOptions ...callopt.Option) (r *education.SetRolePermissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetRolePermission(ctx, req)
}

func (p *kEducationserviceClient) SetRoleMenu(ctx context.Context, req *education.SetRoleMenuReq, callOptions ...callopt.Option) (r *education.SetRoleMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetRoleMenu(ctx, req)
}

func (p *kEducationserviceClient) CreateMenu(ctx context.Context, req *education.CreateMenuReq, callOptions ...callopt.Option) (r *education.CreateMenuResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMenu(ctx, req)
}

func (p *kEducationserviceClient) CreateOffice(ctx context.Context, req *education.CreateOfficeReq, callOptions ...callopt.Option) (r *education.CreateOfficeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateOffice(ctx, req)
}

func (p *kEducationserviceClient) GetOfficeList(ctx context.Context, req *education.GetOfficeListReq, callOptions ...callopt.Option) (r *education.GetOfficeListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOfficeList(ctx, req)
}

func (p *kEducationserviceClient) ImportCourseApply(ctx context.Context, req *education.ImportCourseApplyReq, callOptions ...callopt.Option) (r *education.ImportCourseApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ImportCourseApply(ctx, req)
}

func (p *kEducationserviceClient) GetCourseApplyList(ctx context.Context, req *education.GetCourseApplyListReq, callOptions ...callopt.Option) (r *education.GetCourseApplyListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCourseApplyList(ctx, req)
}

func (p *kEducationserviceClient) GetSelfCourseApplyList(ctx context.Context, req *education.GetSelfCourseApplyListReq, callOptions ...callopt.Option) (r *education.GetSelfCourseApplyListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSelfCourseApplyList(ctx, req)
}

func (p *kEducationserviceClient) DeleteCourseApply(ctx context.Context, req *education.DeleteCourseApplyReq, callOptions ...callopt.Option) (r *education.DeleteCourseApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCourseApply(ctx, req)
}

func (p *kEducationserviceClient) ApproveCourseApply(ctx context.Context, req *education.ApproveCourseApplyReq, callOptions ...callopt.Option) (r *education.ApproveCourseApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ApproveCourseApply(ctx, req)
}

func (p *kEducationserviceClient) ResetCourseApply(ctx context.Context, req *education.ResetCourseApplyReq, callOptions ...callopt.Option) (r *education.ResetCourseApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ResetCourseApply(ctx, req)
}
