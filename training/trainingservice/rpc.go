package trainingservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/training"
)

func ImportInternship(ctx context.Context, req *training.ImportInternshipReq, callOptions ...client.Option) (resp *training.ImportInternshipResp, err error) {
	clients := MustNewClient("training", callOptions...)
	return clients.ImportInternship(ctx, req)
}
