package middleware

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
	"time"
)

type Backchannel interface {
	UserId() string
	UserEnterTime() string
}

type BackchannelUser struct {
	Uid       string
	EnterTime string
}

func (bu BackchannelUser) UserId() string {
	return bu.Uid
}

func (bu BackchannelUser) UserEnterTime() string {
	return bu.EnterTime
}

func BackchannelAuth(gid string) martini.Handler {
	return func(r render.Render, req *http.Request, res http.ResponseWriter, c martini.Context) {
		backchannel := BackchannelUser{Uid: "1", EnterTime: time.Now().Format("2006-01-02 15:04:05")}
		// 每接收一个请求都会new一个新的实例
		c.MapTo(&backchannel, (*Backchannel)(nil))
	}
}
