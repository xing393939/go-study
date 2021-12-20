package controllers

import (
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/cast"
	"iris_demo/models"
	"iris_demo/service"
	"log"
)

type UserController struct {
	Service service.UserService
}

//查询所有/api/user/list
func (con *UserController) GetList() (result *models.Result) {
	return con.Service.GetUserList()
}

//保存and修改/api/user/save/user
func (con *UserController) PostSaveUser(ctx context.Context) (result models.Result) {
	user := new(models.User)
	err := ctx.ReadForm(user)
	if err != nil {
		log.Println(err)
		result.Msg = "数据有错误"
		return
	}
	return con.Service.PostSaveUser(*user)
}

//根据id查询/api/user/user/by/id?id=2
func (con *UserController) GetUserById(ctx context.Context) (result models.Result) {
	id := ctx.URLParam("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return con.Service.GetUserById(cast.ToUint(id))
}

//根据id删除
func (con *UserController) PostDelUser(ctx context.Context) (result models.Result) {
	id := ctx.PostValue("id")
	if id == "" {
		result.Code = 400
		result.Msg = "缺少参数id"
		return
	}
	return con.Service.DelUser(cast.ToUint(id))
}
