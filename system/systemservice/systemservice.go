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
	"GetUserList": kitex.NewMethodInfo(
		getUserListHandler,
		newSystemserviceGetUserListArgs,
		newSystemserviceGetUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAuthCode": kitex.NewMethodInfo(
		getAuthCodeHandler,
		newSystemserviceGetAuthCodeArgs,
		newSystemserviceGetAuthCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMenuList": kitex.NewMethodInfo(
		getMenuListHandler,
		newSystemserviceGetMenuListArgs,
		newSystemserviceGetMenuListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetRoleList": kitex.NewMethodInfo(
		getRoleListHandler,
		newSystemserviceGetRoleListArgs,
		newSystemserviceGetRoleListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateMenu": kitex.NewMethodInfo(
		createMenuHandler,
		newSystemserviceCreateMenuArgs,
		newSystemserviceCreateMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateMenu": kitex.NewMethodInfo(
		updateMenuHandler,
		newSystemserviceUpdateMenuArgs,
		newSystemserviceUpdateMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetApiList": kitex.NewMethodInfo(
		getApiListHandler,
		newSystemserviceGetApiListArgs,
		newSystemserviceGetApiListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteMenu": kitex.NewMethodInfo(
		deleteMenuHandler,
		newSystemserviceDeleteMenuArgs,
		newSystemserviceDeleteMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAppList": kitex.NewMethodInfo(
		getAppListHandler,
		newSystemserviceGetAppListArgs,
		newSystemserviceGetAppListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateApp": kitex.NewMethodInfo(
		createAppHandler,
		newSystemserviceCreateAppArgs,
		newSystemserviceCreateAppResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetApp": kitex.NewMethodInfo(
		getAppHandler,
		newSystemserviceGetAppArgs,
		newSystemserviceGetAppResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateEduUser": kitex.NewMethodInfo(
		createEduUserHandler,
		newSystemserviceCreateEduUserArgs,
		newSystemserviceCreateEduUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEduUserList": kitex.NewMethodInfo(
		getEduUserListHandler,
		newSystemserviceGetEduUserListArgs,
		newSystemserviceGetEduUserListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPassword": kitex.NewMethodInfo(
		getPasswordHandler,
		newSystemserviceGetPasswordArgs,
		newSystemserviceGetPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetAppRoleMenu": kitex.NewMethodInfo(
		setAppRoleMenuHandler,
		newSystemserviceSetAppRoleMenuArgs,
		newSystemserviceSetAppRoleMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetEduMenuList": kitex.NewMethodInfo(
		getEduMenuListHandler,
		newSystemserviceGetEduMenuListArgs,
		newSystemserviceGetEduMenuListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateEduMenu": kitex.NewMethodInfo(
		createEduMenuHandler,
		newSystemserviceCreateEduMenuArgs,
		newSystemserviceCreateEduMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateEduMenu": kitex.NewMethodInfo(
		updateEduMenuHandler,
		newSystemserviceUpdateEduMenuArgs,
		newSystemserviceUpdateEduMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteEduMenu": kitex.NewMethodInfo(
		deleteEduMenuHandler,
		newSystemserviceDeleteEduMenuArgs,
		newSystemserviceDeleteEduMenuResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteApp": kitex.NewMethodInfo(
		deleteAppHandler,
		newSystemserviceDeleteAppArgs,
		newSystemserviceDeleteAppResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetAppStatus": kitex.NewMethodInfo(
		setAppStatusHandler,
		newSystemserviceSetAppStatusArgs,
		newSystemserviceSetAppStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SetEduMenuStatus": kitex.NewMethodInfo(
		setEduMenuStatusHandler,
		newSystemserviceSetEduMenuStatusArgs,
		newSystemserviceSetEduMenuStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetOperationLogsList": kitex.NewMethodInfo(
		getOperationLogsListHandler,
		newSystemserviceGetOperationLogsListArgs,
		newSystemserviceGetOperationLogsListResult,
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

func getUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetUserListArgs)
	realResult := result.(*system.SystemserviceGetUserListResult)
	success, err := handler.(system.Systemservice).GetUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetUserListArgs() interface{} {
	return system.NewSystemserviceGetUserListArgs()
}

func newSystemserviceGetUserListResult() interface{} {
	return system.NewSystemserviceGetUserListResult()
}

func getAuthCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetAuthCodeArgs)
	realResult := result.(*system.SystemserviceGetAuthCodeResult)
	success, err := handler.(system.Systemservice).GetAuthCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetAuthCodeArgs() interface{} {
	return system.NewSystemserviceGetAuthCodeArgs()
}

func newSystemserviceGetAuthCodeResult() interface{} {
	return system.NewSystemserviceGetAuthCodeResult()
}

func getMenuListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetMenuListArgs)
	realResult := result.(*system.SystemserviceGetMenuListResult)
	success, err := handler.(system.Systemservice).GetMenuList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetMenuListArgs() interface{} {
	return system.NewSystemserviceGetMenuListArgs()
}

func newSystemserviceGetMenuListResult() interface{} {
	return system.NewSystemserviceGetMenuListResult()
}

func getRoleListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetRoleListArgs)
	realResult := result.(*system.SystemserviceGetRoleListResult)
	success, err := handler.(system.Systemservice).GetRoleList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetRoleListArgs() interface{} {
	return system.NewSystemserviceGetRoleListArgs()
}

func newSystemserviceGetRoleListResult() interface{} {
	return system.NewSystemserviceGetRoleListResult()
}

func createMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateMenuArgs)
	realResult := result.(*system.SystemserviceCreateMenuResult)
	success, err := handler.(system.Systemservice).CreateMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateMenuArgs() interface{} {
	return system.NewSystemserviceCreateMenuArgs()
}

func newSystemserviceCreateMenuResult() interface{} {
	return system.NewSystemserviceCreateMenuResult()
}

func updateMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceUpdateMenuArgs)
	realResult := result.(*system.SystemserviceUpdateMenuResult)
	success, err := handler.(system.Systemservice).UpdateMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceUpdateMenuArgs() interface{} {
	return system.NewSystemserviceUpdateMenuArgs()
}

func newSystemserviceUpdateMenuResult() interface{} {
	return system.NewSystemserviceUpdateMenuResult()
}

func getApiListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetApiListArgs)
	realResult := result.(*system.SystemserviceGetApiListResult)
	success, err := handler.(system.Systemservice).GetApiList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetApiListArgs() interface{} {
	return system.NewSystemserviceGetApiListArgs()
}

func newSystemserviceGetApiListResult() interface{} {
	return system.NewSystemserviceGetApiListResult()
}

func deleteMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceDeleteMenuArgs)
	realResult := result.(*system.SystemserviceDeleteMenuResult)
	success, err := handler.(system.Systemservice).DeleteMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceDeleteMenuArgs() interface{} {
	return system.NewSystemserviceDeleteMenuArgs()
}

func newSystemserviceDeleteMenuResult() interface{} {
	return system.NewSystemserviceDeleteMenuResult()
}

func getAppListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetAppListArgs)
	realResult := result.(*system.SystemserviceGetAppListResult)
	success, err := handler.(system.Systemservice).GetAppList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetAppListArgs() interface{} {
	return system.NewSystemserviceGetAppListArgs()
}

func newSystemserviceGetAppListResult() interface{} {
	return system.NewSystemserviceGetAppListResult()
}

func createAppHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateAppArgs)
	realResult := result.(*system.SystemserviceCreateAppResult)
	success, err := handler.(system.Systemservice).CreateApp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateAppArgs() interface{} {
	return system.NewSystemserviceCreateAppArgs()
}

func newSystemserviceCreateAppResult() interface{} {
	return system.NewSystemserviceCreateAppResult()
}

func getAppHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetAppArgs)
	realResult := result.(*system.SystemserviceGetAppResult)
	success, err := handler.(system.Systemservice).GetApp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetAppArgs() interface{} {
	return system.NewSystemserviceGetAppArgs()
}

func newSystemserviceGetAppResult() interface{} {
	return system.NewSystemserviceGetAppResult()
}

func createEduUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateEduUserArgs)
	realResult := result.(*system.SystemserviceCreateEduUserResult)
	success, err := handler.(system.Systemservice).CreateEduUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateEduUserArgs() interface{} {
	return system.NewSystemserviceCreateEduUserArgs()
}

func newSystemserviceCreateEduUserResult() interface{} {
	return system.NewSystemserviceCreateEduUserResult()
}

func getEduUserListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetEduUserListArgs)
	realResult := result.(*system.SystemserviceGetEduUserListResult)
	success, err := handler.(system.Systemservice).GetEduUserList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetEduUserListArgs() interface{} {
	return system.NewSystemserviceGetEduUserListArgs()
}

func newSystemserviceGetEduUserListResult() interface{} {
	return system.NewSystemserviceGetEduUserListResult()
}

func getPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetPasswordArgs)
	realResult := result.(*system.SystemserviceGetPasswordResult)
	success, err := handler.(system.Systemservice).GetPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetPasswordArgs() interface{} {
	return system.NewSystemserviceGetPasswordArgs()
}

func newSystemserviceGetPasswordResult() interface{} {
	return system.NewSystemserviceGetPasswordResult()
}

func setAppRoleMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceSetAppRoleMenuArgs)
	realResult := result.(*system.SystemserviceSetAppRoleMenuResult)
	success, err := handler.(system.Systemservice).SetAppRoleMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceSetAppRoleMenuArgs() interface{} {
	return system.NewSystemserviceSetAppRoleMenuArgs()
}

func newSystemserviceSetAppRoleMenuResult() interface{} {
	return system.NewSystemserviceSetAppRoleMenuResult()
}

func getEduMenuListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetEduMenuListArgs)
	realResult := result.(*system.SystemserviceGetEduMenuListResult)
	success, err := handler.(system.Systemservice).GetEduMenuList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetEduMenuListArgs() interface{} {
	return system.NewSystemserviceGetEduMenuListArgs()
}

func newSystemserviceGetEduMenuListResult() interface{} {
	return system.NewSystemserviceGetEduMenuListResult()
}

func createEduMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceCreateEduMenuArgs)
	realResult := result.(*system.SystemserviceCreateEduMenuResult)
	success, err := handler.(system.Systemservice).CreateEduMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceCreateEduMenuArgs() interface{} {
	return system.NewSystemserviceCreateEduMenuArgs()
}

func newSystemserviceCreateEduMenuResult() interface{} {
	return system.NewSystemserviceCreateEduMenuResult()
}

func updateEduMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceUpdateEduMenuArgs)
	realResult := result.(*system.SystemserviceUpdateEduMenuResult)
	success, err := handler.(system.Systemservice).UpdateEduMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceUpdateEduMenuArgs() interface{} {
	return system.NewSystemserviceUpdateEduMenuArgs()
}

func newSystemserviceUpdateEduMenuResult() interface{} {
	return system.NewSystemserviceUpdateEduMenuResult()
}

func deleteEduMenuHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceDeleteEduMenuArgs)
	realResult := result.(*system.SystemserviceDeleteEduMenuResult)
	success, err := handler.(system.Systemservice).DeleteEduMenu(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceDeleteEduMenuArgs() interface{} {
	return system.NewSystemserviceDeleteEduMenuArgs()
}

func newSystemserviceDeleteEduMenuResult() interface{} {
	return system.NewSystemserviceDeleteEduMenuResult()
}

func deleteAppHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceDeleteAppArgs)
	realResult := result.(*system.SystemserviceDeleteAppResult)
	success, err := handler.(system.Systemservice).DeleteApp(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceDeleteAppArgs() interface{} {
	return system.NewSystemserviceDeleteAppArgs()
}

func newSystemserviceDeleteAppResult() interface{} {
	return system.NewSystemserviceDeleteAppResult()
}

func setAppStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceSetAppStatusArgs)
	realResult := result.(*system.SystemserviceSetAppStatusResult)
	success, err := handler.(system.Systemservice).SetAppStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceSetAppStatusArgs() interface{} {
	return system.NewSystemserviceSetAppStatusArgs()
}

func newSystemserviceSetAppStatusResult() interface{} {
	return system.NewSystemserviceSetAppStatusResult()
}

func setEduMenuStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceSetEduMenuStatusArgs)
	realResult := result.(*system.SystemserviceSetEduMenuStatusResult)
	success, err := handler.(system.Systemservice).SetEduMenuStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceSetEduMenuStatusArgs() interface{} {
	return system.NewSystemserviceSetEduMenuStatusArgs()
}

func newSystemserviceSetEduMenuStatusResult() interface{} {
	return system.NewSystemserviceSetEduMenuStatusResult()
}

func getOperationLogsListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*system.SystemserviceGetOperationLogsListArgs)
	realResult := result.(*system.SystemserviceGetOperationLogsListResult)
	success, err := handler.(system.Systemservice).GetOperationLogsList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSystemserviceGetOperationLogsListArgs() interface{} {
	return system.NewSystemserviceGetOperationLogsListArgs()
}

func newSystemserviceGetOperationLogsListResult() interface{} {
	return system.NewSystemserviceGetOperationLogsListResult()
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

func (p *kClient) GetUserList(ctx context.Context, req *system.GetUserListReq) (r *system.GetUserListResp, err error) {
	var _args system.SystemserviceGetUserListArgs
	_args.Req = req
	var _result system.SystemserviceGetUserListResult
	if err = p.c.Call(ctx, "GetUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAuthCode(ctx context.Context, req *system.GetAuthCodeReq) (r *system.GetAuthCodeResp, err error) {
	var _args system.SystemserviceGetAuthCodeArgs
	_args.Req = req
	var _result system.SystemserviceGetAuthCodeResult
	if err = p.c.Call(ctx, "GetAuthCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMenuList(ctx context.Context, req *system.GetMenuListReq) (r *system.GetMenuListResp, err error) {
	var _args system.SystemserviceGetMenuListArgs
	_args.Req = req
	var _result system.SystemserviceGetMenuListResult
	if err = p.c.Call(ctx, "GetMenuList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetRoleList(ctx context.Context, req *system.GetRoleListReq) (r *system.GetRoleListResp, err error) {
	var _args system.SystemserviceGetRoleListArgs
	_args.Req = req
	var _result system.SystemserviceGetRoleListResult
	if err = p.c.Call(ctx, "GetRoleList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateMenu(ctx context.Context, req *system.CreateMenuReq) (r *system.CreateMenuResp, err error) {
	var _args system.SystemserviceCreateMenuArgs
	_args.Req = req
	var _result system.SystemserviceCreateMenuResult
	if err = p.c.Call(ctx, "CreateMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateMenu(ctx context.Context, req *system.UpdateMenuReq) (r *system.UpdateMenuResp, err error) {
	var _args system.SystemserviceUpdateMenuArgs
	_args.Req = req
	var _result system.SystemserviceUpdateMenuResult
	if err = p.c.Call(ctx, "UpdateMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetApiList(ctx context.Context, req *system.GetApiListReq) (r *system.GetApiListResp, err error) {
	var _args system.SystemserviceGetApiListArgs
	_args.Req = req
	var _result system.SystemserviceGetApiListResult
	if err = p.c.Call(ctx, "GetApiList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteMenu(ctx context.Context, req *system.DeleteMenuReq) (r *system.DeleteMenuResp, err error) {
	var _args system.SystemserviceDeleteMenuArgs
	_args.Req = req
	var _result system.SystemserviceDeleteMenuResult
	if err = p.c.Call(ctx, "DeleteMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAppList(ctx context.Context, req *system.GetAppListReq) (r *system.GetAppListResp, err error) {
	var _args system.SystemserviceGetAppListArgs
	_args.Req = req
	var _result system.SystemserviceGetAppListResult
	if err = p.c.Call(ctx, "GetAppList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateApp(ctx context.Context, req *system.CreateAppReq) (r *system.CreateAppResp, err error) {
	var _args system.SystemserviceCreateAppArgs
	_args.Req = req
	var _result system.SystemserviceCreateAppResult
	if err = p.c.Call(ctx, "CreateApp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetApp(ctx context.Context, req *system.GetAppReq) (r *system.GetAppResp, err error) {
	var _args system.SystemserviceGetAppArgs
	_args.Req = req
	var _result system.SystemserviceGetAppResult
	if err = p.c.Call(ctx, "GetApp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateEduUser(ctx context.Context, req *system.CreateEduUserReq) (r *system.CreateEduUserResp, err error) {
	var _args system.SystemserviceCreateEduUserArgs
	_args.Req = req
	var _result system.SystemserviceCreateEduUserResult
	if err = p.c.Call(ctx, "CreateEduUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEduUserList(ctx context.Context, req *system.GetEduUserListReq) (r *system.GetEduUserListResp, err error) {
	var _args system.SystemserviceGetEduUserListArgs
	_args.Req = req
	var _result system.SystemserviceGetEduUserListResult
	if err = p.c.Call(ctx, "GetEduUserList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPassword(ctx context.Context, req *system.GetPasswordReq) (r *system.GetPasswordResp, err error) {
	var _args system.SystemserviceGetPasswordArgs
	_args.Req = req
	var _result system.SystemserviceGetPasswordResult
	if err = p.c.Call(ctx, "GetPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetAppRoleMenu(ctx context.Context, req *system.SetAppRoleMenuReq) (r *system.SetAppRoleMenuResp, err error) {
	var _args system.SystemserviceSetAppRoleMenuArgs
	_args.Req = req
	var _result system.SystemserviceSetAppRoleMenuResult
	if err = p.c.Call(ctx, "SetAppRoleMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetEduMenuList(ctx context.Context, req *system.GetEduMenuListReq) (r *system.GetEduMenuListResp, err error) {
	var _args system.SystemserviceGetEduMenuListArgs
	_args.Req = req
	var _result system.SystemserviceGetEduMenuListResult
	if err = p.c.Call(ctx, "GetEduMenuList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateEduMenu(ctx context.Context, req *system.CreateEduMenuReq) (r *system.CreateEduMenuResp, err error) {
	var _args system.SystemserviceCreateEduMenuArgs
	_args.Req = req
	var _result system.SystemserviceCreateEduMenuResult
	if err = p.c.Call(ctx, "CreateEduMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateEduMenu(ctx context.Context, req *system.UpdateEduMenuReq) (r *system.UpdateEduMenuResp, err error) {
	var _args system.SystemserviceUpdateEduMenuArgs
	_args.Req = req
	var _result system.SystemserviceUpdateEduMenuResult
	if err = p.c.Call(ctx, "UpdateEduMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteEduMenu(ctx context.Context, req *system.DeleteEduMenuReq) (r *system.DeleteEduMenuResp, err error) {
	var _args system.SystemserviceDeleteEduMenuArgs
	_args.Req = req
	var _result system.SystemserviceDeleteEduMenuResult
	if err = p.c.Call(ctx, "DeleteEduMenu", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteApp(ctx context.Context, req *system.DeleteAppReq) (r *system.DeleteAppResp, err error) {
	var _args system.SystemserviceDeleteAppArgs
	_args.Req = req
	var _result system.SystemserviceDeleteAppResult
	if err = p.c.Call(ctx, "DeleteApp", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetAppStatus(ctx context.Context, req *system.SetAppStatusReq) (r *system.SetAppStatusResp, err error) {
	var _args system.SystemserviceSetAppStatusArgs
	_args.Req = req
	var _result system.SystemserviceSetAppStatusResult
	if err = p.c.Call(ctx, "SetAppStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetEduMenuStatus(ctx context.Context, req *system.SetEduMenuStatusReq) (r *system.SetEduMenuStatusResp, err error) {
	var _args system.SystemserviceSetEduMenuStatusArgs
	_args.Req = req
	var _result system.SystemserviceSetEduMenuStatusResult
	if err = p.c.Call(ctx, "SetEduMenuStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetOperationLogsList(ctx context.Context, req *system.GetOperationLogsListReq) (r *system.GetOperationLogsListResp, err error) {
	var _args system.SystemserviceGetOperationLogsListArgs
	_args.Req = req
	var _result system.SystemserviceGetOperationLogsListResult
	if err = p.c.Call(ctx, "GetOperationLogsList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
