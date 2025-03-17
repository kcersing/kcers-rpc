// Code generated by Kitex v0.9.1. DO NOT EDIT.

package companyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "rpc_gen/kitex_gen/base"
	contract "rpc_gen/kitex_gen/company/contract"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ContractList(ctx context.Context, req *contract.ContractListReq, callOptions ...callopt.Option) (r *contract.ContractListResp, err error)
	ContractCreate(ctx context.Context, req *contract.CreateOrUpdateContractReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ContractUpdate(ctx context.Context, req *contract.CreateOrUpdateContractReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ContractUpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ContractByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *contract.ContractInfo, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCompanyServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCompanyServiceClient struct {
	*kClient
}

func (p *kCompanyServiceClient) ContractList(ctx context.Context, req *contract.ContractListReq, callOptions ...callopt.Option) (r *contract.ContractListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContractList(ctx, req)
}

func (p *kCompanyServiceClient) ContractCreate(ctx context.Context, req *contract.CreateOrUpdateContractReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContractCreate(ctx, req)
}

func (p *kCompanyServiceClient) ContractUpdate(ctx context.Context, req *contract.CreateOrUpdateContractReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContractUpdate(ctx, req)
}

func (p *kCompanyServiceClient) ContractUpdateStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContractUpdateStatus(ctx, req)
}

func (p *kCompanyServiceClient) ContractByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *contract.ContractInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContractByID(ctx, req)
}
