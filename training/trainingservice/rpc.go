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
