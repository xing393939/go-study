package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go-study/gf002/internal/model"
)

type UserListReq struct {
	g.Meta      `path:"/users" method:"get" tags:"用户" summary:"获取用户列表" dc:"获取用户列表"`
	ContentType string `json:"contentType" dc:"当传递空时表示获取所有"`
}

type UserListRes struct {
	RowsWithUser []model.RowWithUser
	Rows         []model.Row
}
