package http

import (
	"bbs/app/http/module/qa"
	"bbs/app/http/module/user"
	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
	ginSwagger "github.com/gohade/hade/framework/middleware/gin-swagger"
	"github.com/gohade/hade/framework/middleware/gin-swagger/swaggerFiles"
)

// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 用户模块
	user.RegisterRoutes(r)
	// 问答模块
	qa.RegisterRoutes(r)
}
