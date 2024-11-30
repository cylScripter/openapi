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
