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
