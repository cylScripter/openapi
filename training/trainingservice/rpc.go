package trainingservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/training"
)

func ImportTrainingCourse(ctx context.Context, req *training.ImportTrainingCourseReq, callOptions ...client.Option) (resp *training.ImportTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ImportTrainingCourse(ctx, req)
}

func CreateTrainingCourse(ctx context.Context, req *training.CreateTrainingCourseReq, callOptions ...client.Option) (resp *training.CreateTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.CreateTrainingCourse(ctx, req)
}

func GetTrainingCourseList(ctx context.Context, req *training.GetTrainingCourseListReq, callOptions ...client.Option) (resp *training.GetTrainingCourseListResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetTrainingCourseList(ctx, req)
}

func GetSelfTrainingCourseList(ctx context.Context, req *training.GetSelfTrainingCourseListReq, callOptions ...client.Option) (resp *training.GetSelfTrainingCourseListResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetSelfTrainingCourseList(ctx, req)
}

func GetTrainingCourse(ctx context.Context, req *training.GetTrainingCourseReq, callOptions ...client.Option) (resp *training.GetTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetTrainingCourse(ctx, req)
}

func FillTrainingCourse(ctx context.Context, req *training.FillTrainingCourseReq, callOptions ...client.Option) (resp *training.FillTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.FillTrainingCourse(ctx, req)
}

func DeleteTrainingCourse(ctx context.Context, req *training.DeleteTrainingCourseReq, callOptions ...client.Option) (resp *training.DeleteTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.DeleteTrainingCourse(ctx, req)
}

func ExportTrainingCourse(ctx context.Context, req *training.ExportTrainingCourseReq, callOptions ...client.Option) (resp *training.ExportTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ExportTrainingCourse(ctx, req)
}

func UpdateTrainingCourse(ctx context.Context, req *training.UpdateTrainingCourseReq, callOptions ...client.Option) (resp *training.UpdateTrainingCourseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UpdateTrainingCourse(ctx, req)
}

func UpdateTrainingCourseTeacher(ctx context.Context, req *training.UpdateTrainingCourseTeacherReq, callOptions ...client.Option) (resp *training.UpdateTrainingCourseTeacherResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UpdateTrainingCourseTeacher(ctx, req)
}

func ApproveTrainingCourseData(ctx context.Context, req *training.ApproveTrainingCourseDataReq, callOptions ...client.Option) (resp *training.ApproveTrainingCourseDataResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ApproveTrainingCourseData(ctx, req)
}

func UploadTrainingCourseFile(ctx context.Context, req *training.UploadTrainingCourseFileReq, callOptions ...client.Option) (resp *training.UploadTrainingCourseFileResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UploadTrainingCourseFile(ctx, req)
}

func UploadTrainingCourseCase(ctx context.Context, req *training.UploadTrainingCourseCaseReq, callOptions ...client.Option) (resp *training.UploadTrainingCourseCaseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UploadTrainingCourseCase(ctx, req)
}

func ExportTrainingCourseFile(ctx context.Context, req *training.ExportTrainingCourseFileReq, callOptions ...client.Option) (resp *training.ExportTrainingCourseFileResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ExportTrainingCourseFile(ctx, req)
}

func ExportTrainingCourseCase(ctx context.Context, req *training.ExportTrainingCourseCaseReq, callOptions ...client.Option) (resp *training.ExportTrainingCourseCaseResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ExportTrainingCourseCase(ctx, req)
}

func UploadTrainingCourseSource(ctx context.Context, req *training.UploadTrainingCourseSourceReq, callOptions ...client.Option) (resp *training.UploadTrainingCourseSourceResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UploadTrainingCourseSource(ctx, req)
}

func ExportTrainingCourseSource(ctx context.Context, req *training.ExportTrainingCourseSourceReq, callOptions ...client.Option) (resp *training.ExportTrainingCourseSourceResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ExportTrainingCourseSource(ctx, req)
}

func GetExportResult_(ctx context.Context, req *training.GetExportResultReq, callOptions ...client.Option) (resp *training.GetExportResultResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetExportResult_(ctx, req)
}

func GetPastMajorList(ctx context.Context, req *training.GetPastMajorListReq, callOptions ...client.Option) (resp *training.GetPastMajorListResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetPastMajorList(ctx, req)
}

func DeletePastMajor(ctx context.Context, req *training.DeletePastMajorReq, callOptions ...client.Option) (resp *training.DeletePastMajorResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.DeletePastMajor(ctx, req)
}

func AddPastMajor(ctx context.Context, req *training.AddPastMajorReq, callOptions ...client.Option) (resp *training.AddPastMajorResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.AddPastMajor(ctx, req)
}

func UpdatePastMajor(ctx context.Context, req *training.UpdatePastMajorReq, callOptions ...client.Option) (resp *training.UpdatePastMajorResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.UpdatePastMajor(ctx, req)
}

func GetClassroomList(ctx context.Context, req *training.GetClassroomListReq, callOptions ...client.Option) (resp *training.GetClassroomListResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetClassroomList(ctx, req)
}

func ImportClassroom(ctx context.Context, req *training.ImportClassroomReq, callOptions ...client.Option) (resp *training.ImportClassroomResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ImportClassroom(ctx, req)
}

func EditClassroom(ctx context.Context, req *training.EditClassroomReq, callOptions ...client.Option) (resp *training.EditClassroomResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.EditClassroom(ctx, req)
}

func ExportMajorPracticeTeaching(ctx context.Context, req *training.ExportMajorPracticeTeachingReq, callOptions ...client.Option) (resp *training.ExportMajorPracticeTeachingResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ExportMajorPracticeTeaching(ctx, req)
}

func ImportMajorPracticeTeaching(ctx context.Context, req *training.ImportMajorPracticeTeachingReq, callOptions ...client.Option) (resp *training.ImportMajorPracticeTeachingResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ImportMajorPracticeTeaching(ctx, req)
}

func GetMajorPracticeTeachingPlanList(ctx context.Context, req *training.GetMajorPracticeTeachingPlanListReq, callOptions ...client.Option) (resp *training.GetMajorPracticeTeachingPlanListResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetMajorPracticeTeachingPlanList(ctx, req)
}

func FillMajorPracticeTeachingBudget(ctx context.Context, req *training.FillMajorPracticeTeachingBudgetReq, callOptions ...client.Option) (resp *training.FillMajorPracticeTeachingBudgetResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.FillMajorPracticeTeachingBudget(ctx, req)
}

func EditMajorPracticeTeachingPlan(ctx context.Context, req *training.EditMajorPracticeTeachingPlanReq, callOptions ...client.Option) (resp *training.EditMajorPracticeTeachingPlanResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.EditMajorPracticeTeachingPlan(ctx, req)
}

func ScheduleTeaching(ctx context.Context, req *training.ScheduleTeachingReq, callOptions ...client.Option) (resp *training.ScheduleTeachingResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ScheduleTeaching(ctx, req)
}

func ScheduleClassroom(ctx context.Context, req *training.ScheduleClassroomReq, callOptions ...client.Option) (resp *training.ScheduleClassroomResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ScheduleClassroom(ctx, req)
}

func GetTeachSchedule(ctx context.Context, req *training.GetTeachScheduleReq, callOptions ...client.Option) (resp *training.GetTeachScheduleResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetTeachSchedule(ctx, req)
}

func GetClassroomUsage(ctx context.Context, req *training.GetClassroomUsageReq, callOptions ...client.Option) (resp *training.GetClassroomUsageResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetClassroomUsage(ctx, req)
}

func GetPlanDetail(ctx context.Context, req *training.GetPlanDetailReq, callOptions ...client.Option) (resp *training.GetPlanDetailResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.GetPlanDetail(ctx, req)
}
