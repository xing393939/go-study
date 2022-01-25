package handler

import (
	"context"
	"gf001/internal/service"

	"gf001/apiv1"
)

var (
	User = hUser{}
)

type hUser struct{}

func (h *hUser) Hello(ctx context.Context, req *apiv1.UserListReq) (res *apiv1.UserListRes, err error) {
	res = &apiv1.UserListRes{}
	res.RowsWithUser = service.User().GetRowsWithUser(ctx)
	res.Rows = service.User().GetRows(ctx)
	return
}
