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

func GetMenuList(ctx context.Context, req *education.GetMenuListReq, callOptions ...client.Option) (resp *education.GetMenuListResp, err error) {
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

func GetClassList(ctx context.Context, req *education.GetClassListReq, callOptions ...client.Option) (resp *education.GetClassListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetClassList(ctx, req)
}

func ImportCourseApply(ctx context.Context, req *education.ImportCourseApplyReq, callOptions ...client.Option) (resp *education.ImportCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportCourseApply(ctx, req)
}

func GetCourseApplyList(ctx context.Context, req *education.GetCourseApplyListReq, callOptions ...client.Option) (resp *education.GetCourseApplyListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetCourseApplyList(ctx, req)
}

func GetSelfCourseApplyList(ctx context.Context, req *education.GetSelfCourseApplyListReq, callOptions ...client.Option) (resp *education.GetSelfCourseApplyListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetSelfCourseApplyList(ctx, req)
}

func DeleteCourseApply(ctx context.Context, req *education.DeleteCourseApplyReq, callOptions ...client.Option) (resp *education.DeleteCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteCourseApply(ctx, req)
}

func ApproveCourseApply(ctx context.Context, req *education.ApproveCourseApplyReq, callOptions ...client.Option) (resp *education.ApproveCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ApproveCourseApply(ctx, req)
}

func ResetCourseApply(ctx context.Context, req *education.ResetCourseApplyReq, callOptions ...client.Option) (resp *education.ResetCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ResetCourseApply(ctx, req)
}

func FillCourseApply(ctx context.Context, req *education.FillCourseApplyReq, callOptions ...client.Option) (resp *education.FillCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.FillCourseApply(ctx, req)
}

func ExportSelfCheckTable(ctx context.Context, req *education.ExportSelfCheckTableReq, callOptions ...client.Option) (resp *education.ExportSelfCheckTableResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportSelfCheckTable(ctx, req)
}

func ExportResponsibility(ctx context.Context, req *education.ExportResponsibilityReq, callOptions ...client.Option) (resp *education.ExportResponsibilityResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportResponsibility(ctx, req)
}

func ExportAllSelfCheckTable(ctx context.Context, req *education.ExportAllSelfCheckTableReq, callOptions ...client.Option) (resp *education.ExportAllSelfCheckTableResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportAllSelfCheckTable(ctx, req)
}

func UpdateCourseApply(ctx context.Context, req *education.UpdateCourseApplyReq, callOptions ...client.Option) (resp *education.UpdateCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateCourseApply(ctx, req)
}

func ExportCourseApply(ctx context.Context, req *education.ExportCourseApplyReq, callOptions ...client.Option) (resp *education.ExportCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportCourseApply(ctx, req)
}

func ExportCourseApplySummary(ctx context.Context, req *education.ExportCourseApplySummaryReq, callOptions ...client.Option) (resp *education.ExportCourseApplySummaryResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportCourseApplySummary(ctx, req)
}

func OneKeyApproveCourseApply(ctx context.Context, req *education.OneKeyApproveCourseApplyReq, callOptions ...client.Option) (resp *education.OneKeyApproveCourseApplyResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.OneKeyApproveCourseApply(ctx, req)
}

func GetSelectData(ctx context.Context, req *education.GetSelectDataReq, callOptions ...client.Option) (resp *education.GetSelectDataResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetSelectData(ctx, req)
}
