package controllers

import (
	"bytes"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"runtime"
	"strconv"
	"time"
)

type ShareForPerRequest struct {
	ExString string
	Gid      uint64
}

type ShareForAllRequest struct {
	ExString string
	Gid      uint64
}

func GetGoroutineId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func MyDo(perCtx *ShareForPerRequest, global *ShareForAllRequest, r render.Render) {
	r.JSON(200, map[string]interface{}{
		"myGid":  GetGoroutineId(),
		"per":    perCtx,
		"global": global,
	})
}

func MyHandle1(c martini.Context) {
	reqC := &ShareForPerRequest{
		ExString: "perRequest " + time.Now().Format("2006-01-02 15:04:05"),
		Gid:      GetGoroutineId(),
	}
	c.Map(reqC)
}

func MyHandle2(reqCtx *ShareForPerRequest) string {
	return reqCtx.ExString
}

func MyHandle3(global *ShareForAllRequest, r render.Render) {
	r.JSON(200, map[string]interface{}{
		"myGid":  GetGoroutineId(),
		"global": global,
	})
}
