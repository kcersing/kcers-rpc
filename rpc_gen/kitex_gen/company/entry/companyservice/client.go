// Code generated by Kitex v0.9.1. DO NOT EDIT.

package companyservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	base "rpc_gen/kitex_gen/base"
	entry "rpc_gen/kitex_gen/company/entry"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateEntry(ctx context.Context, req *entry.CreateEntry, callOptions ...callopt.Option) (r *base.NilResponse, err error)
	EntryList(ctx context.Context, req *entry.EntryListReq, callOptions ...callopt.Option) (r *entry.EntryListResp, err error)
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

func (p *kCompanyServiceClient) CreateEntry(ctx context.Context, req *entry.CreateEntry, callOptions ...callopt.Option) (r *base.NilResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateEntry(ctx, req)
}

func (p *kCompanyServiceClient) EntryList(ctx context.Context, req *entry.EntryListReq, callOptions ...callopt.Option) (r *entry.EntryListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EntryList(ctx, req)
}
