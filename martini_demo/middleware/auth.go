package middleware

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

type Backchannel interface {
	UserId() string
	GuestId() string
}

type BackchannelUser struct {
	Uid string
	Gid string
}

func (bu BackchannelUser) UserId() string {
	return bu.Uid
}

func (bu BackchannelUser) GuestId() string {
	return bu.Gid
}

func BackchannelAuth(gid string) martini.Handler {
	return func(r render.Render, req *http.Request, res http.ResponseWriter, c martini.Context) {
		backchannel := BackchannelUser{Uid: "1", Gid: gid}
		c.MapTo(&backchannel, (*Backchannel)(nil))
	}
}
