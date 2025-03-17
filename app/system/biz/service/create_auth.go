package service

import (
	"context"
	base "rpc_gen/kitex_gen/base"
	auth "rpc_gen/kitex_gen/system/auth"
	"strconv"
	"system/biz/infra/service"
)

type CreateAuthService struct {
	ctx context.Context
} // NewCreateAuthService new CreateAuthService
func NewCreateAuthService(ctx context.Context) *CreateAuthService {
	return &CreateAuthService{ctx: ctx}
}

// Run create note info
func (s *CreateAuthService) Run(req *auth.CreateOrUpdateApiAuthReq) (resp *base.NilResponse, err error) {
	// Finish your business logic.

	err = service.NewAuth(s.ctx).UpdateApiAuth(strconv.FormatInt(req.RoleId, 10), req.Apis)
	if err != nil {
		return nil, err
	}
	return
}
