package service

import (
	"context"
	user "rpc_gen/kitex_gen/user"
	"testing"
)

func TestTokenList_Run(t *testing.T) {
	ctx := context.Background()
	s := NewTokenListService(ctx)
	// init req and assert value

	req := &user.TokenListReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
