// Code generated by Kitex v0.9.1. DO NOT EDIT.

package scheduleservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "rpc_gen/kitex_gen/base"
	schedule "rpc_gen/kitex_gen/schedule"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateSchedule": kitex.NewMethodInfo(
		createScheduleHandler,
		newScheduleServiceCreateScheduleArgs,
		newScheduleServiceCreateScheduleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateSchedule": kitex.NewMethodInfo(
		updateScheduleHandler,
		newScheduleServiceUpdateScheduleArgs,
		newScheduleServiceUpdateScheduleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateStatus": kitex.NewMethodInfo(
		updateStatusHandler,
		newScheduleServiceUpdateStatusArgs,
		newScheduleServiceUpdateStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ScheduleList": kitex.NewMethodInfo(
		scheduleListHandler,
		newScheduleServiceScheduleListArgs,
		newScheduleServiceScheduleListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ScheduleDateList": kitex.NewMethodInfo(
		scheduleDateListHandler,
		newScheduleServiceScheduleDateListArgs,
		newScheduleServiceScheduleDateListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ScheduleInfo": kitex.NewMethodInfo(
		scheduleInfoHandler,
		newScheduleServiceScheduleInfoArgs,
		newScheduleServiceScheduleInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"MemberList": kitex.NewMethodInfo(
		memberListHandler,
		newScheduleServiceMemberListArgs,
		newScheduleServiceMemberListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CreateMember": kitex.NewMethodInfo(
		createMemberHandler,
		newScheduleServiceCreateMemberArgs,
		newScheduleServiceCreateMemberResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateMemberStatus": kitex.NewMethodInfo(
		updateMemberStatusHandler,
		newScheduleServiceUpdateMemberStatusArgs,
		newScheduleServiceUpdateMemberStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"SearchSubscribeByMember": kitex.NewMethodInfo(
		searchSubscribeByMemberHandler,
		newScheduleServiceSearchSubscribeByMemberArgs,
		newScheduleServiceSearchSubscribeByMemberResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"CoachList": kitex.NewMethodInfo(
		coachListHandler,
		newScheduleServiceCoachListArgs,
		newScheduleServiceCoachListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateCoachStatus": kitex.NewMethodInfo(
		updateCoachStatusHandler,
		newScheduleServiceUpdateCoachStatusArgs,
		newScheduleServiceUpdateCoachStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	scheduleServiceServiceInfo                = NewServiceInfo()
	scheduleServiceServiceInfoForClient       = NewServiceInfoForClient()
	scheduleServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return scheduleServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return scheduleServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return scheduleServiceServiceInfoForClient
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
	serviceName := "ScheduleService"
	handlerType := (*schedule.ScheduleService)(nil)
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
		"PackageName": "schedule",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceCreateScheduleArgs)
	realResult := result.(*schedule.ScheduleServiceCreateScheduleResult)
	success, err := handler.(schedule.ScheduleService).CreateSchedule(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceCreateScheduleArgs() interface{} {
	return schedule.NewScheduleServiceCreateScheduleArgs()
}

func newScheduleServiceCreateScheduleResult() interface{} {
	return schedule.NewScheduleServiceCreateScheduleResult()
}

func updateScheduleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceUpdateScheduleArgs)
	realResult := result.(*schedule.ScheduleServiceUpdateScheduleResult)
	success, err := handler.(schedule.ScheduleService).UpdateSchedule(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceUpdateScheduleArgs() interface{} {
	return schedule.NewScheduleServiceUpdateScheduleArgs()
}

func newScheduleServiceUpdateScheduleResult() interface{} {
	return schedule.NewScheduleServiceUpdateScheduleResult()
}

func updateStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceUpdateStatusArgs)
	realResult := result.(*schedule.ScheduleServiceUpdateStatusResult)
	success, err := handler.(schedule.ScheduleService).UpdateStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceUpdateStatusArgs() interface{} {
	return schedule.NewScheduleServiceUpdateStatusArgs()
}

func newScheduleServiceUpdateStatusResult() interface{} {
	return schedule.NewScheduleServiceUpdateStatusResult()
}

func scheduleListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceScheduleListArgs)
	realResult := result.(*schedule.ScheduleServiceScheduleListResult)
	success, err := handler.(schedule.ScheduleService).ScheduleList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceScheduleListArgs() interface{} {
	return schedule.NewScheduleServiceScheduleListArgs()
}

func newScheduleServiceScheduleListResult() interface{} {
	return schedule.NewScheduleServiceScheduleListResult()
}

func scheduleDateListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceScheduleDateListArgs)
	realResult := result.(*schedule.ScheduleServiceScheduleDateListResult)
	success, err := handler.(schedule.ScheduleService).ScheduleDateList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceScheduleDateListArgs() interface{} {
	return schedule.NewScheduleServiceScheduleDateListArgs()
}

func newScheduleServiceScheduleDateListResult() interface{} {
	return schedule.NewScheduleServiceScheduleDateListResult()
}

func scheduleInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceScheduleInfoArgs)
	realResult := result.(*schedule.ScheduleServiceScheduleInfoResult)
	success, err := handler.(schedule.ScheduleService).ScheduleInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceScheduleInfoArgs() interface{} {
	return schedule.NewScheduleServiceScheduleInfoArgs()
}

func newScheduleServiceScheduleInfoResult() interface{} {
	return schedule.NewScheduleServiceScheduleInfoResult()
}

func memberListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceMemberListArgs)
	realResult := result.(*schedule.ScheduleServiceMemberListResult)
	success, err := handler.(schedule.ScheduleService).MemberList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceMemberListArgs() interface{} {
	return schedule.NewScheduleServiceMemberListArgs()
}

func newScheduleServiceMemberListResult() interface{} {
	return schedule.NewScheduleServiceMemberListResult()
}

func createMemberHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceCreateMemberArgs)
	realResult := result.(*schedule.ScheduleServiceCreateMemberResult)
	success, err := handler.(schedule.ScheduleService).CreateMember(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceCreateMemberArgs() interface{} {
	return schedule.NewScheduleServiceCreateMemberArgs()
}

func newScheduleServiceCreateMemberResult() interface{} {
	return schedule.NewScheduleServiceCreateMemberResult()
}

func updateMemberStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceUpdateMemberStatusArgs)
	realResult := result.(*schedule.ScheduleServiceUpdateMemberStatusResult)
	success, err := handler.(schedule.ScheduleService).UpdateMemberStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceUpdateMemberStatusArgs() interface{} {
	return schedule.NewScheduleServiceUpdateMemberStatusArgs()
}

func newScheduleServiceUpdateMemberStatusResult() interface{} {
	return schedule.NewScheduleServiceUpdateMemberStatusResult()
}

func searchSubscribeByMemberHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceSearchSubscribeByMemberArgs)
	realResult := result.(*schedule.ScheduleServiceSearchSubscribeByMemberResult)
	success, err := handler.(schedule.ScheduleService).SearchSubscribeByMember(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceSearchSubscribeByMemberArgs() interface{} {
	return schedule.NewScheduleServiceSearchSubscribeByMemberArgs()
}

func newScheduleServiceSearchSubscribeByMemberResult() interface{} {
	return schedule.NewScheduleServiceSearchSubscribeByMemberResult()
}

func coachListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceCoachListArgs)
	realResult := result.(*schedule.ScheduleServiceCoachListResult)
	success, err := handler.(schedule.ScheduleService).CoachList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceCoachListArgs() interface{} {
	return schedule.NewScheduleServiceCoachListArgs()
}

func newScheduleServiceCoachListResult() interface{} {
	return schedule.NewScheduleServiceCoachListResult()
}

func updateCoachStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*schedule.ScheduleServiceUpdateCoachStatusArgs)
	realResult := result.(*schedule.ScheduleServiceUpdateCoachStatusResult)
	success, err := handler.(schedule.ScheduleService).UpdateCoachStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newScheduleServiceUpdateCoachStatusArgs() interface{} {
	return schedule.NewScheduleServiceUpdateCoachStatusArgs()
}

func newScheduleServiceUpdateCoachStatusResult() interface{} {
	return schedule.NewScheduleServiceUpdateCoachStatusResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq) (r *base.NilResponse, err error) {
	var _args schedule.ScheduleServiceCreateScheduleArgs
	_args.Req = req
	var _result schedule.ScheduleServiceCreateScheduleResult
	if err = p.c.Call(ctx, "CreateSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateSchedule(ctx context.Context, req *schedule.CreateOrUpdateScheduleReq) (r *base.NilResponse, err error) {
	var _args schedule.ScheduleServiceUpdateScheduleArgs
	_args.Req = req
	var _result schedule.ScheduleServiceUpdateScheduleResult
	if err = p.c.Call(ctx, "UpdateSchedule", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args schedule.ScheduleServiceUpdateStatusArgs
	_args.Req = req
	var _result schedule.ScheduleServiceUpdateStatusResult
	if err = p.c.Call(ctx, "UpdateStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ScheduleList(ctx context.Context, req *schedule.ScheduleListReq) (r *schedule.ScheduleListResp, err error) {
	var _args schedule.ScheduleServiceScheduleListArgs
	_args.Req = req
	var _result schedule.ScheduleServiceScheduleListResult
	if err = p.c.Call(ctx, "ScheduleList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ScheduleDateList(ctx context.Context, req *schedule.ScheduleListReq) (r *schedule.ScheduleDateListResp, err error) {
	var _args schedule.ScheduleServiceScheduleDateListArgs
	_args.Req = req
	var _result schedule.ScheduleServiceScheduleDateListResult
	if err = p.c.Call(ctx, "ScheduleDateList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ScheduleInfo(ctx context.Context, req *base.IDReq) (r *schedule.ScheduleInfo, err error) {
	var _args schedule.ScheduleServiceScheduleInfoArgs
	_args.Req = req
	var _result schedule.ScheduleServiceScheduleInfoResult
	if err = p.c.Call(ctx, "ScheduleInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MemberList(ctx context.Context, req *schedule.ScheduleMemberListReq) (r *schedule.ScheduleMemberListResp, err error) {
	var _args schedule.ScheduleServiceMemberListArgs
	_args.Req = req
	var _result schedule.ScheduleServiceMemberListResult
	if err = p.c.Call(ctx, "MemberList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateMember(ctx context.Context, req *schedule.CreateMemberReq) (r *schedule.ScheduleMemberListResp, err error) {
	var _args schedule.ScheduleServiceCreateMemberArgs
	_args.Req = req
	var _result schedule.ScheduleServiceCreateMemberResult
	if err = p.c.Call(ctx, "CreateMember", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateMemberStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args schedule.ScheduleServiceUpdateMemberStatusArgs
	_args.Req = req
	var _result schedule.ScheduleServiceUpdateMemberStatusResult
	if err = p.c.Call(ctx, "UpdateMemberStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchSubscribeByMember(ctx context.Context, req *schedule.SearchSubscribeByMemberReq) (r *schedule.SearchSubscribeByMemberResp, err error) {
	var _args schedule.ScheduleServiceSearchSubscribeByMemberArgs
	_args.Req = req
	var _result schedule.ScheduleServiceSearchSubscribeByMemberResult
	if err = p.c.Call(ctx, "SearchSubscribeByMember", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CoachList(ctx context.Context, req *schedule.ScheduleCoachListReq) (r *schedule.ScheduleCoachListResp, err error) {
	var _args schedule.ScheduleServiceCoachListArgs
	_args.Req = req
	var _result schedule.ScheduleServiceCoachListResult
	if err = p.c.Call(ctx, "CoachList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateCoachStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args schedule.ScheduleServiceUpdateCoachStatusArgs
	_args.Req = req
	var _result schedule.ScheduleServiceUpdateCoachStatusResult
	if err = p.c.Call(ctx, "UpdateCoachStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
