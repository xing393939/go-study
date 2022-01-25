package service

import (
	"context"
	"gf001/internal/model"
	"gf001/internal/service/internal/dao"
)

type sUser struct {}

func User() *sUser {
	return &sUser{}
}

func (s *sUser) GetRowsWithUser(ctx context.Context) (out []model.RowWithUser) {
	m := dao.Users.Ctx(ctx)
	all,_ := m.All()
	all.ScanList(&out, "User")
	return
}

func (s *sUser) GetRows(ctx context.Context) (out []model.Row) {
	m := dao.Users.Ctx(ctx)
	_ = m.Scan(&out)
	return
}
