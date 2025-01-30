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

func GetUserInfo(ctx context.Context, req *education.GetUserInfoReq, callOptions ...client.Option) (resp *education.GetUserInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetUserInfo(ctx, req)
}

func GetAuthCode(ctx context.Context, req *education.GetAuthCodeReq, callOptions ...client.Option) (resp *education.GetAuthCodeResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetAuthCode(ctx, req)
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

func ImportUser(ctx context.Context, req *education.ImportUserReq, callOptions ...client.Option) (resp *education.ImportUserResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportUser(ctx, req)
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

func GetTeacherInfoList(ctx context.Context, req *education.GetTeacherInfoListReq, callOptions ...client.Option) (resp *education.GetTeacherInfoListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetTeacherInfoList(ctx, req)
}

func ImportTeacherInfo(ctx context.Context, req *education.ImportTeacherInfoReq, callOptions ...client.Option) (resp *education.ImportTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportTeacherInfo(ctx, req)
}

func UpdateTeacherInfo(ctx context.Context, req *education.UpdateTeacherInfoReq, callOptions ...client.Option) (resp *education.UpdateTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateTeacherInfo(ctx, req)
}

func SetTeacherInfoStatus(ctx context.Context, req *education.SetTeacherInfoStatusReq, callOptions ...client.Option) (resp *education.SetTeacherInfoStatusResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SetTeacherInfoStatus(ctx, req)
}

func OderTeacherInfo(ctx context.Context, req *education.OrderTeacherInfoReq, callOptions ...client.Option) (resp *education.OrderTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.OderTeacherInfo(ctx, req)
}

func UpdateMenu(ctx context.Context, req *education.UpdateMenuReq, callOptions ...client.Option) (resp *education.UpdateMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateMenu(ctx, req)
}

func DeleteMenu(ctx context.Context, req *education.DeleteMenuReq, callOptions ...client.Option) (resp *education.DeleteMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteMenu(ctx, req)
}

func GetToken(ctx context.Context, req *education.GetTokenReq, callOptions ...client.Option) (resp *education.GetTokenResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetToken(ctx, req)
}

func UpdateRole(ctx context.Context, req *education.UpdateRoleReq, callOptions ...client.Option) (resp *education.UpdateRoleResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateRole(ctx, req)
}

func SetRoleStatus(ctx context.Context, req *education.SetRoleStatusReq, callOptions ...client.Option) (resp *education.SetRoleStatusResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SetRoleStatus(ctx, req)
}

func DeleteRole(ctx context.Context, req *education.DeleteRoleReq, callOptions ...client.Option) (resp *education.DeleteRoleResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteRole(ctx, req)
}

func DeleteOffice(ctx context.Context, req *education.DeleteOfficeReq, callOptions ...client.Option) (resp *education.DeleteOfficeResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteOffice(ctx, req)
}

func UpdateOffice(ctx context.Context, req *education.UpdateOfficeReq, callOptions ...client.Option) (resp *education.UpdateOfficeResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateOffice(ctx, req)
}

func UpdateUser(ctx context.Context, req *education.UpdateUserReq, callOptions ...client.Option) (resp *education.UpdateUserResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateUser(ctx, req)
}

func ResetPassword(ctx context.Context, req *education.ResetPasswordReq, callOptions ...client.Option) (resp *education.ResetPasswordResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ResetPassword(ctx, req)
}

func DeleteTeacherInfo(ctx context.Context, req *education.DeleteTeacherInfoReq, callOptions ...client.Option) (resp *education.DeleteTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteTeacherInfo(ctx, req)
}

func CreateTeacherInfo(ctx context.Context, req *education.CreateTeacherInfoReq, callOptions ...client.Option) (resp *education.CreateTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateTeacherInfo(ctx, req)
}

func ExportTeacherInfo(ctx context.Context, req *education.ExportTeacherInfoReq, callOptions ...client.Option) (resp *education.ExportTeacherInfoResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportTeacherInfo(ctx, req)
}

func GetCourseList(ctx context.Context, req *education.GetCourseListReq, callOptions ...client.Option) (resp *education.GetCourseListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetCourseList(ctx, req)
}

func DeleteCourse(ctx context.Context, req *education.DeleteCourseReq, callOptions ...client.Option) (resp *education.DeleteCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteCourse(ctx, req)
}

func ImportCourse(ctx context.Context, req *education.ImportCourseReq, callOptions ...client.Option) (resp *education.ImportCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportCourse(ctx, req)
}

func UpdateCourse(ctx context.Context, req *education.UpdateCourseReq, callOptions ...client.Option) (resp *education.UpdateCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateCourse(ctx, req)
}
