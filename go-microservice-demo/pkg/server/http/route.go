package http

import (
	"fmt"
	"github.com/win5do/golang-microservice-demo/pkg/api/petpb"
	"github.com/win5do/golang-microservice-demo/pkg/repository/db/dbcore"
	"net/http"

	"github.com/gin-gonic/gin"
	petdb "github.com/win5do/golang-microservice-demo/pkg/repository/db/pet"
	petsvc "github.com/win5do/golang-microservice-demo/pkg/service/pet"
)

func Register(mux *gin.Engine) {
	// list all api
	mux.GET("/apis", func(c *gin.Context) {
		list := ""
		for _, v := range mux.Routes() {
			list += fmt.Sprintf("%s %s\n", v.Method, v.Path)
		}
		c.String(http.StatusOK, list)
	})

	svc := petsvc.NewPetService(dbcore.NewTxImpl(), petdb.NewPetDomain())
	mux.GET("/api/v1/ping", func(c *gin.Context) {
		req := petpb.Id{}
		req.Id = "1"
		res, _ := svc.Ping(c, &req)
		c.JSON(200, res)
	})
}
