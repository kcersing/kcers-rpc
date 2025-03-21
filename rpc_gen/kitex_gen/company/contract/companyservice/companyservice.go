// Code generated by Kitex v0.9.1. DO NOT EDIT.

package companyservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ContractList": kitex.NewMethodInfo(
		contractListHandler,
		newCompanyServiceContractListArgs,
		newCompanyServiceContractListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ContractCreate": kitex.NewMethodInfo(
		contractCreateHandler,
		newCompanyServiceContractCreateArgs,
		newCompanyServiceContractCreateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ContractUpdate": kitex.NewMethodInfo(
		contractUpdateHandler,
		newCompanyServiceContractUpdateArgs,
		newCompanyServiceContractUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ContractUpdateStatus": kitex.NewMethodInfo(
		contractUpdateStatusHandler,
		newCompanyServiceContractUpdateStatusArgs,
		newCompanyServiceContractUpdateStatusResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"ContractByID": kitex.NewMethodInfo(
		contractByIDHandler,
		newCompanyServiceContractByIDArgs,
		newCompanyServiceContractByIDResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	companyServiceServiceInfo                = NewServiceInfo()
	companyServiceServiceInfoForClient       = NewServiceInfoForClient()
	companyServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return companyServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return companyServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return companyServiceServiceInfoForClient
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
	serviceName := "CompanyService"
	handlerType := (*contract.CompanyService)(nil)
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
		"PackageName": "contract",
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

func contractListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contract.CompanyServiceContractListArgs)
	realResult := result.(*contract.CompanyServiceContractListResult)
	success, err := handler.(contract.CompanyService).ContractList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompanyServiceContractListArgs() interface{} {
	return contract.NewCompanyServiceContractListArgs()
}

func newCompanyServiceContractListResult() interface{} {
	return contract.NewCompanyServiceContractListResult()
}

func contractCreateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contract.CompanyServiceContractCreateArgs)
	realResult := result.(*contract.CompanyServiceContractCreateResult)
	success, err := handler.(contract.CompanyService).ContractCreate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompanyServiceContractCreateArgs() interface{} {
	return contract.NewCompanyServiceContractCreateArgs()
}

func newCompanyServiceContractCreateResult() interface{} {
	return contract.NewCompanyServiceContractCreateResult()
}

func contractUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contract.CompanyServiceContractUpdateArgs)
	realResult := result.(*contract.CompanyServiceContractUpdateResult)
	success, err := handler.(contract.CompanyService).ContractUpdate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompanyServiceContractUpdateArgs() interface{} {
	return contract.NewCompanyServiceContractUpdateArgs()
}

func newCompanyServiceContractUpdateResult() interface{} {
	return contract.NewCompanyServiceContractUpdateResult()
}

func contractUpdateStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contract.CompanyServiceContractUpdateStatusArgs)
	realResult := result.(*contract.CompanyServiceContractUpdateStatusResult)
	success, err := handler.(contract.CompanyService).ContractUpdateStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompanyServiceContractUpdateStatusArgs() interface{} {
	return contract.NewCompanyServiceContractUpdateStatusArgs()
}

func newCompanyServiceContractUpdateStatusResult() interface{} {
	return contract.NewCompanyServiceContractUpdateStatusResult()
}

func contractByIDHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*contract.CompanyServiceContractByIDArgs)
	realResult := result.(*contract.CompanyServiceContractByIDResult)
	success, err := handler.(contract.CompanyService).ContractByID(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCompanyServiceContractByIDArgs() interface{} {
	return contract.NewCompanyServiceContractByIDArgs()
}

func newCompanyServiceContractByIDResult() interface{} {
	return contract.NewCompanyServiceContractByIDResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ContractList(ctx context.Context, req *contract.ContractListReq) (r *contract.ContractListResp, err error) {
	var _args contract.CompanyServiceContractListArgs
	_args.Req = req
	var _result contract.CompanyServiceContractListResult
	if err = p.c.Call(ctx, "ContractList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContractCreate(ctx context.Context, req *contract.CreateOrUpdateContractReq) (r *base.NilResponse, err error) {
	var _args contract.CompanyServiceContractCreateArgs
	_args.Req = req
	var _result contract.CompanyServiceContractCreateResult
	if err = p.c.Call(ctx, "ContractCreate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContractUpdate(ctx context.Context, req *contract.CreateOrUpdateContractReq) (r *base.NilResponse, err error) {
	var _args contract.CompanyServiceContractUpdateArgs
	_args.Req = req
	var _result contract.CompanyServiceContractUpdateResult
	if err = p.c.Call(ctx, "ContractUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContractUpdateStatus(ctx context.Context, req *base.StatusCodeReq) (r *base.NilResponse, err error) {
	var _args contract.CompanyServiceContractUpdateStatusArgs
	_args.Req = req
	var _result contract.CompanyServiceContractUpdateStatusResult
	if err = p.c.Call(ctx, "ContractUpdateStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ContractByID(ctx context.Context, req *base.IDReq) (r *contract.ContractInfo, err error) {
	var _args contract.CompanyServiceContractByIDArgs
	_args.Req = req
	var _result contract.CompanyServiceContractByIDResult
	if err = p.c.Call(ctx, "ContractByID", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
