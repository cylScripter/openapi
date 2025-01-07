// Code generated by Kitex v0.11.3. DO NOT EDIT.

package commonservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	common "github.com/cylScripter/openapi/common"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"UploadFile": kitex.NewMethodInfo(
		uploadFileHandler,
		newCommonserviceUploadFileArgs,
		newCommonserviceUploadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CompleteFile": kitex.NewMethodInfo(
		completeFileHandler,
		newCommonserviceCompleteFileArgs,
		newCommonserviceCompleteFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadNewMultipart": kitex.NewMethodInfo(
		uploadNewMultipartHandler,
		newCommonserviceUploadNewMultipartArgs,
		newCommonserviceUploadNewMultipartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetPresignedUrlList": kitex.NewMethodInfo(
		getPresignedUrlListHandler,
		newCommonserviceGetPresignedUrlListArgs,
		newCommonserviceGetPresignedUrlListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CompleteMultipart": kitex.NewMethodInfo(
		completeMultipartHandler,
		newCommonserviceCompleteMultipartArgs,
		newCommonserviceCompleteMultipartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"AbortMultipart": kitex.NewMethodInfo(
		abortMultipartHandler,
		newCommonserviceAbortMultipartArgs,
		newCommonserviceAbortMultipartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetObject": kitex.NewMethodInfo(
		getObjectHandler,
		newCommonserviceGetObjectArgs,
		newCommonserviceGetObjectResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteObject": kitex.NewMethodInfo(
		deleteObjectHandler,
		newCommonserviceDeleteObjectArgs,
		newCommonserviceDeleteObjectResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateAsyncTask": kitex.NewMethodInfo(
		createAsyncTaskHandler,
		newCommonserviceCreateAsyncTaskArgs,
		newCommonserviceCreateAsyncTaskResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetAsyncTaskResult": kitex.NewMethodInfo(
		getAsyncTaskResult_Handler,
		newCommonserviceGetAsyncTaskResultArgs,
		newCommonserviceGetAsyncTaskResultResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	commonserviceServiceInfo                = NewServiceInfo()
	commonserviceServiceInfoForClient       = NewServiceInfoForClient()
	commonserviceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commonserviceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commonserviceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commonserviceServiceInfoForClient
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
	serviceName := "commonservice"
	handlerType := (*common.Commonservice)(nil)
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
		"PackageName": "common",
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

func uploadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceUploadFileArgs)
	realResult := result.(*common.CommonserviceUploadFileResult)
	success, err := handler.(common.Commonservice).UploadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceUploadFileArgs() interface{} {
	return common.NewCommonserviceUploadFileArgs()
}

func newCommonserviceUploadFileResult() interface{} {
	return common.NewCommonserviceUploadFileResult()
}

func completeFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceCompleteFileArgs)
	realResult := result.(*common.CommonserviceCompleteFileResult)
	success, err := handler.(common.Commonservice).CompleteFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceCompleteFileArgs() interface{} {
	return common.NewCommonserviceCompleteFileArgs()
}

func newCommonserviceCompleteFileResult() interface{} {
	return common.NewCommonserviceCompleteFileResult()
}

func uploadNewMultipartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceUploadNewMultipartArgs)
	realResult := result.(*common.CommonserviceUploadNewMultipartResult)
	success, err := handler.(common.Commonservice).UploadNewMultipart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceUploadNewMultipartArgs() interface{} {
	return common.NewCommonserviceUploadNewMultipartArgs()
}

func newCommonserviceUploadNewMultipartResult() interface{} {
	return common.NewCommonserviceUploadNewMultipartResult()
}

func getPresignedUrlListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceGetPresignedUrlListArgs)
	realResult := result.(*common.CommonserviceGetPresignedUrlListResult)
	success, err := handler.(common.Commonservice).GetPresignedUrlList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceGetPresignedUrlListArgs() interface{} {
	return common.NewCommonserviceGetPresignedUrlListArgs()
}

func newCommonserviceGetPresignedUrlListResult() interface{} {
	return common.NewCommonserviceGetPresignedUrlListResult()
}

func completeMultipartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceCompleteMultipartArgs)
	realResult := result.(*common.CommonserviceCompleteMultipartResult)
	success, err := handler.(common.Commonservice).CompleteMultipart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceCompleteMultipartArgs() interface{} {
	return common.NewCommonserviceCompleteMultipartArgs()
}

func newCommonserviceCompleteMultipartResult() interface{} {
	return common.NewCommonserviceCompleteMultipartResult()
}

func abortMultipartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceAbortMultipartArgs)
	realResult := result.(*common.CommonserviceAbortMultipartResult)
	success, err := handler.(common.Commonservice).AbortMultipart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceAbortMultipartArgs() interface{} {
	return common.NewCommonserviceAbortMultipartArgs()
}

func newCommonserviceAbortMultipartResult() interface{} {
	return common.NewCommonserviceAbortMultipartResult()
}

func getObjectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceGetObjectArgs)
	realResult := result.(*common.CommonserviceGetObjectResult)
	success, err := handler.(common.Commonservice).GetObject(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceGetObjectArgs() interface{} {
	return common.NewCommonserviceGetObjectArgs()
}

func newCommonserviceGetObjectResult() interface{} {
	return common.NewCommonserviceGetObjectResult()
}

func deleteObjectHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceDeleteObjectArgs)
	realResult := result.(*common.CommonserviceDeleteObjectResult)
	success, err := handler.(common.Commonservice).DeleteObject(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceDeleteObjectArgs() interface{} {
	return common.NewCommonserviceDeleteObjectArgs()
}

func newCommonserviceDeleteObjectResult() interface{} {
	return common.NewCommonserviceDeleteObjectResult()
}

func createAsyncTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceCreateAsyncTaskArgs)
	realResult := result.(*common.CommonserviceCreateAsyncTaskResult)
	success, err := handler.(common.Commonservice).CreateAsyncTask(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceCreateAsyncTaskArgs() interface{} {
	return common.NewCommonserviceCreateAsyncTaskArgs()
}

func newCommonserviceCreateAsyncTaskResult() interface{} {
	return common.NewCommonserviceCreateAsyncTaskResult()
}

func getAsyncTaskResult_Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*common.CommonserviceGetAsyncTaskResultArgs)
	realResult := result.(*common.CommonserviceGetAsyncTaskResultResult)
	success, err := handler.(common.Commonservice).GetAsyncTaskResult_(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommonserviceGetAsyncTaskResultArgs() interface{} {
	return common.NewCommonserviceGetAsyncTaskResultArgs()
}

func newCommonserviceGetAsyncTaskResultResult() interface{} {
	return common.NewCommonserviceGetAsyncTaskResultResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UploadFile(ctx context.Context, req *common.UploadFileReq) (r *common.UploadFileResp, err error) {
	var _args common.CommonserviceUploadFileArgs
	_args.Req = req
	var _result common.CommonserviceUploadFileResult
	if err = p.c.Call(ctx, "UploadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CompleteFile(ctx context.Context, req *common.CompleteFileReq) (r *common.CompleteFileResp, err error) {
	var _args common.CommonserviceCompleteFileArgs
	_args.Req = req
	var _result common.CommonserviceCompleteFileResult
	if err = p.c.Call(ctx, "CompleteFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadNewMultipart(ctx context.Context, req *common.UploadNewMultipartReq) (r *common.UploadNewMultipartResp, err error) {
	var _args common.CommonserviceUploadNewMultipartArgs
	_args.Req = req
	var _result common.CommonserviceUploadNewMultipartResult
	if err = p.c.Call(ctx, "UploadNewMultipart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPresignedUrlList(ctx context.Context, req *common.GetPresignedUrlListReq) (r *common.GetPresignedUrlListResp, err error) {
	var _args common.CommonserviceGetPresignedUrlListArgs
	_args.Req = req
	var _result common.CommonserviceGetPresignedUrlListResult
	if err = p.c.Call(ctx, "GetPresignedUrlList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CompleteMultipart(ctx context.Context, req *common.CompleteMultipartReq) (r *common.CompleteMultipartResp, err error) {
	var _args common.CommonserviceCompleteMultipartArgs
	_args.Req = req
	var _result common.CommonserviceCompleteMultipartResult
	if err = p.c.Call(ctx, "CompleteMultipart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AbortMultipart(ctx context.Context, req *common.AbortMultipartReq) (r *common.AbortMultipartResp, err error) {
	var _args common.CommonserviceAbortMultipartArgs
	_args.Req = req
	var _result common.CommonserviceAbortMultipartResult
	if err = p.c.Call(ctx, "AbortMultipart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetObject(ctx context.Context, req *common.GetObjectReq) (r *common.GetObjectResp, err error) {
	var _args common.CommonserviceGetObjectArgs
	_args.Req = req
	var _result common.CommonserviceGetObjectResult
	if err = p.c.Call(ctx, "GetObject", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteObject(ctx context.Context, req *common.DeleteObjectReq) (r *common.DeleteObjectResp, err error) {
	var _args common.CommonserviceDeleteObjectArgs
	_args.Req = req
	var _result common.CommonserviceDeleteObjectResult
	if err = p.c.Call(ctx, "DeleteObject", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAsyncTask(ctx context.Context, req *common.CreateAsyncTaskReq) (r *common.CreateAsyncTaskResp, err error) {
	var _args common.CommonserviceCreateAsyncTaskArgs
	_args.Req = req
	var _result common.CommonserviceCreateAsyncTaskResult
	if err = p.c.Call(ctx, "CreateAsyncTask", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAsyncTaskResult_(ctx context.Context, req *common.GetAsyncTaskResultReq) (r *common.GetAsyncTaskResultResp, err error) {
	var _args common.CommonserviceGetAsyncTaskResultArgs
	_args.Req = req
	var _result common.CommonserviceGetAsyncTaskResultResult
	if err = p.c.Call(ctx, "GetAsyncTaskResult", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
