// Code generated by Kitex v0.9.1. DO NOT EDIT.

package systemservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *auth.RoleInfo, err error)
	CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *auth.RoleListResp, err error)
	UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	ApiAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*auth.ApiAuthInfo, err error)
	GetLogsList(ctx context.Context, req *auth.LogsListReq, callOptions ...callopt.Option) (r *auth.LogsListReq, err error)
	DeleteLogs(ctx context.Context, req *base.Ids, callOptions ...callopt.Option) (r *base.NilResponse, err error)
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
	return &kSystemServiceClient{
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

type kSystemServiceClient struct {
	*kClient
}

func (p *kSystemServiceClient) CreateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateRole(ctx, req)
}

func (p *kSystemServiceClient) UpdateRole(ctx context.Context, req *auth.RoleInfo, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRole(ctx, req)
}

func (p *kSystemServiceClient) DeleteRole(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRole(ctx, req)
}

func (p *kSystemServiceClient) RoleByID(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r *auth.RoleInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RoleByID(ctx, req)
}

func (p *kSystemServiceClient) CreateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMenuAuth(ctx, req)
}

func (p *kSystemServiceClient) UpdateMenuAuth(ctx context.Context, req *auth.MenuAuthInfoReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMenuAuth(ctx, req)
}

func (p *kSystemServiceClient) RoleList(ctx context.Context, req *base.PageInfoReq, callOptions ...callopt.Option) (r *auth.RoleListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RoleList(ctx, req)
}

func (p *kSystemServiceClient) UpdateRoleStatus(ctx context.Context, req *base.StatusCodeReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRoleStatus(ctx, req)
}

func (p *kSystemServiceClient) CreateAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateAuth(ctx, req)
}

func (p *kSystemServiceClient) UpdateApiAuth(ctx context.Context, req *auth.CreateOrUpdateApiAuthReq, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateApiAuth(ctx, req)
}

func (p *kSystemServiceClient) ApiAuth(ctx context.Context, req *base.IDReq, callOptions ...callopt.Option) (r []*auth.ApiAuthInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ApiAuth(ctx, req)
}

func (p *kSystemServiceClient) GetLogsList(ctx context.Context, req *auth.LogsListReq, callOptions ...callopt.Option) (r *auth.LogsListReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLogsList(ctx, req)
}

func (p *kSystemServiceClient) DeleteLogs(ctx context.Context, req *base.Ids, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLogs(ctx, req)
}
