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
