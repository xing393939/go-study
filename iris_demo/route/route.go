package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_demo/controllers"
	"iris_demo/datasource"
	"iris_demo/repo"
	"iris_demo/service"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"

	db := datasource.GetDB()
	uRepo := repo.NewUserRepository(db)
	userService := service.NewUserService(uRepo)
	mvc.New(app.Party(bathUrl + "/user")).Register(userService).Handle(new(controllers.UserController))

	mRepo := repo.NewMovieRepository(datasource.Movies)
	movieService := service.NewMovieService(mRepo)
	mvc.New(app.Party(bathUrl + "/movies")).Register(movieService).Handle(new(controllers.MovieController))
}
