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

func GetSelfCourseList(ctx context.Context, req *education.GetSelfCourseListReq, callOptions ...client.Option) (resp *education.GetSelfCourseListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetSelfCourseList(ctx, req)
}

func GetCourseApplicationList(ctx context.Context, req *education.GetCourseApplicationListReq, callOptions ...client.Option) (resp *education.GetCourseApplicationListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetCourseApplicationList(ctx, req)
}

func DeleteCourseApplication(ctx context.Context, req *education.DeleteCourseApplicationReq, callOptions ...client.Option) (resp *education.DeleteCourseApplicationResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteCourseApplication(ctx, req)
}

func GetSelfCourseApplicationList(ctx context.Context, req *education.GetSelfCourseApplicationListReq, callOptions ...client.Option) (resp *education.GetSelfCourseApplicationListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetSelfCourseApplicationList(ctx, req)
}

func AdjustCourseApplication(ctx context.Context, req *education.AdjustCourseApplicationReq, callOptions ...client.Option) (resp *education.AdjustCourseApplicationResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.AdjustCourseApplication(ctx, req)
}

func GetHolidayList(ctx context.Context, req *education.GetHolidayListReq, callOptions ...client.Option) (resp *education.GetHolidayListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetHolidayList(ctx, req)
}

func DeleteHoliday(ctx context.Context, req *education.DeleteHolidayReq, callOptions ...client.Option) (resp *education.DeleteHolidayResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteHoliday(ctx, req)
}

func CreateHoliday(ctx context.Context, req *education.CreateHolidayReq, callOptions ...client.Option) (resp *education.CreateHolidayResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateHoliday(ctx, req)
}

func UpdateHoliday(ctx context.Context, req *education.UpdateHolidayReq, callOptions ...client.Option) (resp *education.UpdateHolidayResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateHoliday(ctx, req)
}

func GetInternshipList(ctx context.Context, req *education.GetInternshipListReq, callOptions ...client.Option) (resp *education.GetInternshipListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetInternshipList(ctx, req)
}

func DeleteInternship(ctx context.Context, req *education.DeleteInternshipReq, callOptions ...client.Option) (resp *education.DeleteInternshipResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteInternship(ctx, req)
}

func UpdateInternship(ctx context.Context, req *education.UpdateInternshipReq, callOptions ...client.Option) (resp *education.UpdateInternshipResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateInternship(ctx, req)
}

func CreateInternship(ctx context.Context, req *education.CreateInternshipReq, callOptions ...client.Option) (resp *education.CreateInternshipResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateInternship(ctx, req)
}

func ImportInternship(ctx context.Context, req *education.ImportInternshipReq, callOptions ...client.Option) (resp *education.ImportInternshipResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportInternship(ctx, req)
}

func GetTrainingCourseList(ctx context.Context, req *education.GetTrainingCourseListReq, callOptions ...client.Option) (resp *education.GetTrainingCourseListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetTrainingCourseList(ctx, req)
}

func CreateTrainingCourse(ctx context.Context, req *education.CreateTrainingCourseReq, callOptions ...client.Option) (resp *education.CreateTrainingCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateTrainingCourse(ctx, req)
}

func DeleteTrainingCourse(ctx context.Context, req *education.DeleteTrainingCourseReq, callOptions ...client.Option) (resp *education.DeleteTrainingCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteTrainingCourse(ctx, req)
}

func UpdateTrainingCourse(ctx context.Context, req *education.UpdateTrainingCourseReq, callOptions ...client.Option) (resp *education.UpdateTrainingCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateTrainingCourse(ctx, req)
}

func UpdateTrainingCourseTeacher(ctx context.Context, req *education.UpdateTrainingCourseTeacherReq, callOptions ...client.Option) (resp *education.UpdateTrainingCourseTeacherResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateTrainingCourseTeacher(ctx, req)
}

func ImportTrainingCourse(ctx context.Context, req *education.ImportTrainingCourseReq, callOptions ...client.Option) (resp *education.ImportTrainingCourseResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportTrainingCourse(ctx, req)
}

func GetSelfTrainingCourseList(ctx context.Context, req *education.GetSelfTrainingCourseListReq, callOptions ...client.Option) (resp *education.GetSelfTrainingCourseListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetSelfTrainingCourseList(ctx, req)
}

func SyncFinalExam(ctx context.Context, req *education.SyncFinalExamReq, callOptions ...client.Option) (resp *education.SyncFinalExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.SyncFinalExam(ctx, req)
}

func GetFinalExamList(ctx context.Context, req *education.GetFinalExamListReq, callOptions ...client.Option) (resp *education.GetFinalExamListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetFinalExamList(ctx, req)
}

func GetFinalExamSubjectList(ctx context.Context, req *education.GetFinalExamSubjectListReq, callOptions ...client.Option) (resp *education.GetFinalExamSubjectListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetFinalExamSubjectList(ctx, req)
}

func UpdateFinalExam(ctx context.Context, req *education.UpdateFinalExamReq, callOptions ...client.Option) (resp *education.UpdateFinalExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateFinalExam(ctx, req)
}

func ExportFinalExam(ctx context.Context, req *education.ExportFinalExamReq, callOptions ...client.Option) (resp *education.ExportFinalExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ExportFinalExam(ctx, req)
}

func CalculateFinalExamWorkload(ctx context.Context, req *education.CalculateFinalExamWorkloadReq, callOptions ...client.Option) (resp *education.CalculateFinalExamWorkloadResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CalculateFinalExamWorkload(ctx, req)
}

func DeleteFinalExam(ctx context.Context, req *education.DeleteFinalExamReq, callOptions ...client.Option) (resp *education.DeleteFinalExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteFinalExam(ctx, req)
}

func FillFinalExamPaper(ctx context.Context, req *education.FillFinalExamPaperReq, callOptions ...client.Option) (resp *education.FillFinalExamPaperResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.FillFinalExamPaper(ctx, req)
}

func GetFinalExamFillRecordList(ctx context.Context, req *education.GetFinalExamRecordListReq, callOptions ...client.Option) (resp *education.GetFinalExamRecordListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetFinalExamFillRecordList(ctx, req)
}

func DeleteFinalExamRecord(ctx context.Context, req *education.DeleteFinalExamRecordReq, callOptions ...client.Option) (resp *education.DeleteFinalExamRecordResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteFinalExamRecord(ctx, req)
}

func UpdateFinalExamRecord(ctx context.Context, req *education.UpdateFinalExamRecordReq, callOptions ...client.Option) (resp *education.UpdateFinalExamRecordResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateFinalExamRecord(ctx, req)
}

func GetFinalExamClassList(ctx context.Context, req *education.GetFinalExamClassListReq, callOptions ...client.Option) (resp *education.GetFinalExamClassListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetFinalExamClassList(ctx, req)
}

func GetBeginExamList(ctx context.Context, req *education.GetBeginExamListReq, callOptions ...client.Option) (resp *education.GetBeginExamListResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetBeginExamList(ctx, req)
}

func UpdateBeginExam(ctx context.Context, req *education.UpdateBeginExamReq, callOptions ...client.Option) (resp *education.UpdateBeginExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.UpdateBeginExam(ctx, req)
}

func CreateBeginExam(ctx context.Context, req *education.CreateBeginExamReq, callOptions ...client.Option) (resp *education.CreateBeginExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.CreateBeginExam(ctx, req)
}

func DeleteBeginExam(ctx context.Context, req *education.DeleteBeginExamReq, callOptions ...client.Option) (resp *education.DeleteBeginExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.DeleteBeginExam(ctx, req)
}

func ImportBeginExam(ctx context.Context, req *education.ImportBeginExamReq, callOptions ...client.Option) (resp *education.ImportBeginExamResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.ImportBeginExam(ctx, req)
}

func GetRoleMenu(ctx context.Context, req *education.GetRoleMenuReq, callOptions ...client.Option) (resp *education.GetRoleMenuResp, err error) {
	clients := MustNewClient("education", callOptions...)
	return clients.GetRoleMenu(ctx, req)
}
