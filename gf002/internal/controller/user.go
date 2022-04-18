package controller

import (
	"context"
	"go-study/gf002/internal/service"
	"go-study/gf002/api/v1"
)

var (
	User = hUser{}
)

type hUser struct{}

func (h *hUser) Hello(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res = &v1.UserListRes{}
	res.RowsWithUser = service.User().GetRowsWithUser(ctx)
	res.Rows = service.User().GetRows(ctx)
	return
}
