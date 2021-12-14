package controllers

import (
	"github.com/go-martini/martini"
	"net/http"
	"time"
)

type ShareForPerRequest struct {
	ExString string
	Req      *http.Request
}

type ShareForAllRequest struct {
	ExString string
	Req      *http.Request
}

func MyDo(perCtx *ShareForPerRequest, global *ShareForAllRequest) string {
	return perCtx.ExString + "\n<br>" + global.ExString
}

func MyHandle1(req *http.Request, c martini.Context) {
	reqC := &ShareForPerRequest{
		ExString: "perRequest " + time.Now().Format("2006-01-02 15:04:05"),
		Req:      req,
	}
	c.Map(reqC)
}

func MyHandle2(reqCtx *ShareForPerRequest) string {
	return reqCtx.ExString
}

func MyHandle3(reqCtx *ShareForAllRequest) string {
	return reqCtx.ExString
}
