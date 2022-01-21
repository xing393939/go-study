package http

import (
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/win5do/go-lib/logx"
	"github.com/win5do/golang-microservice-demo/pkg/config"
)

func Run(cfg *config.Config) {
	mux := SetupMux(cfg)
	addr := net.JoinHostPort("", cfg.HttpPort)

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Infof("http server start: %s", addr)
	_ = server.ListenAndServe()
}

func SetupMux(cfg *config.Config) http.Handler {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin
	mux := gin.Default()
	mux.Use(ginhttp.Middleware(cfg.Tracer))

	Register(mux)

	return mux
}
