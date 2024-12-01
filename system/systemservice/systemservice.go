// Code generated by Kitex v0.11.3. DO NOT EDIT.

package systemservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	system "github.com/cylScripter/openapi/system"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newSystemserviceLoginArgs,
		newSystemserviceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateUser": kitex.NewMethodInfo(
		createUserHandler,
		newSystemserviceCreateUserArgs,
		newSystemserviceCreateUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateRole": kitex.NewMethodInfo(
		createRoleHandler,
		newSystemserviceCreateRoleArgs,
		newSystemserviceCreateRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreatePermission": kitex.NewMethodInfo(
		createPermissionHandler,
		newSystemserviceCreatePermissionArgs,
		newSystemserviceCreatePermissionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetUserRole": kitex.NewMethodInfo(
		setUserRoleHandler,
		newSystemserviceSetUserRoleArgs,
		newSystemserviceSetUserRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetRolePermission": kitex.NewMethodInfo(
		setRolePermissionHandler,
		newSystemserviceSetRolePermissionArgs,
		newSystemserviceSetRolePermissionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMenu": kitex.NewMethodInfo(
		getMenuHandler,
		newSystemserviceGetMenuArgs,
		newSystemserviceGetMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newSystemserviceGetUserInfoArgs,
		newSystemserviceGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	systemserviceServiceInfo                = NewServiceInfo()
	systemserviceServiceInfoForClient       = NewServiceInfoForClient()
	systemserviceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return systemserviceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return systemserviceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return systemserviceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "systemservice"
	handlerType := (*system.Systemservice)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "system",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceLoginArgs)
	realResult := result.(*system.SystemserviceLoginResult)
	success, err := handler.(system.Systemservice).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceLoginArgs() interface{} {
	return system.NewSystemserviceLoginArgs()
}

func newSystemserviceLoginResult() interface{} {
	return system.NewSystemserviceLoginResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateUserArgs)
	realResult := result.(*system.SystemserviceCreateUserResult)
	success, err := handler.(system.Systemservice).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateUserArgs() interface{} {
	return system.NewSystemserviceCreateUserArgs()
}

func newSystemserviceCreateUserResult() interface{} {
	return system.NewSystemserviceCreateUserResult()
}

func createRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateRoleArgs)
	realResult := result.(*system.SystemserviceCreateRoleResult)
	success, err := handler.(system.Systemservice).CreateRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateRoleArgs() interface{} {
	return system.NewSystemserviceCreateRoleArgs()
}

func newSystemserviceCreateRoleResult() interface{} {
	return system.NewSystemserviceCreateRoleResult()
}

func createPermissionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreatePermissionArgs)
	realResult := result.(*system.SystemserviceCreatePermissionResult)
	success, err := handler.(system.Systemservice).CreatePermission(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreatePermissionArgs() interface{} {
	return system.NewSystemserviceCreatePermissionArgs()
}

func newSystemserviceCreatePermissionResult() interface{} {
	return system.NewSystemserviceCreatePermissionResult()
}

func setUserRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceSetUserRoleArgs)
	realResult := result.(*system.SystemserviceSetUserRoleResult)
	success, err := handler.(system.Systemservice).SetUserRole(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceSetUserRoleArgs() interface{} {
	return system.NewSystemserviceSetUserRoleArgs()
}

func newSystemserviceSetUserRoleResult() interface{} {
	return system.NewSystemserviceSetUserRoleResult()
}

func setRolePermissionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceSetRolePermissionArgs)
	realResult := result.(*system.SystemserviceSetRolePermissionResult)
	success, err := handler.(system.Systemservice).SetRolePermission(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceSetRolePermissionArgs() interface{} {
	return system.NewSystemserviceSetRolePermissionArgs()
}

func newSystemserviceSetRolePermissionResult() interface{} {
	return system.NewSystemserviceSetRolePermissionResult()
}

func getMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetMenuArgs)
	realResult := result.(*system.SystemserviceGetMenuResult)
	success, err := handler.(system.Systemservice).GetMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetMenuArgs() interface{} {
	return system.NewSystemserviceGetMenuArgs()
}

func newSystemserviceGetMenuResult() interface{} {
	return system.NewSystemserviceGetMenuResult()
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetUserInfoArgs)
	realResult := result.(*system.SystemserviceGetUserInfoResult)
	success, err := handler.(system.Systemservice).GetUserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetUserInfoArgs() interface{} {
	return system.NewSystemserviceGetUserInfoArgs()
}

func newSystemserviceGetUserInfoResult() interface{} {
	return system.NewSystemserviceGetUserInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *system.LoginReq) (r *system.LoginResp, err error) {
	var _args system.SystemserviceLoginArgs
	_args.Req = req
	var _result system.SystemserviceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *system.CreateUserReq) (r *system.CreateUserResp, err error) {
	var _args system.SystemserviceCreateUserArgs
	_args.Req = req
	var _result system.SystemserviceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateRole(ctx context.Context, req *system.CreateRoleReq) (r *system.CreateRoleResp, err error) {
	var _args system.SystemserviceCreateRoleArgs
	_args.Req = req
	var _result system.SystemserviceCreateRoleResult
	if err = p.c.Call(ctx, "CreateRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreatePermission(ctx context.Context, req *system.CreatePermissionReq) (r *system.CreatePermissionResp, err error) {
	var _args system.SystemserviceCreatePermissionArgs
	_args.Req = req
	var _result system.SystemserviceCreatePermissionResult
	if err = p.c.Call(ctx, "CreatePermission", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetUserRole(ctx context.Context, req *system.SetUserRoleReq) (r *system.SetUserRoleResp, err error) {
	var _args system.SystemserviceSetUserRoleArgs
	_args.Req = req
	var _result system.SystemserviceSetUserRoleResult
	if err = p.c.Call(ctx, "SetUserRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetRolePermission(ctx context.Context, req *system.SetRolePermissionReq) (r *system.SetRolePermissionResp, err error) {
	var _args system.SystemserviceSetRolePermissionArgs
	_args.Req = req
	var _result system.SystemserviceSetRolePermissionResult
	if err = p.c.Call(ctx, "SetRolePermission", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMenu(ctx context.Context, req *system.GetMenuReq) (r *system.GetMenuResp, err error) {
	var _args system.SystemserviceGetMenuArgs
	_args.Req = req
	var _result system.SystemserviceGetMenuResult
	if err = p.c.Call(ctx, "GetMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, req *system.GetUserInfoReq) (r *system.GetUserInfoResp, err error) {
	var _args system.SystemserviceGetUserInfoArgs
	_args.Req = req
	var _result system.SystemserviceGetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
