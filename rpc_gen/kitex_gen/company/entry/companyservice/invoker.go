// Code generated by Kitex v0.9.1. DO NOT EDIT.

package companyservice

import (
	server "github.com/cloudwego/kitex/server"
	entry "rpc_gen/kitex_gen/company/entry"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler entry.CompanyService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
