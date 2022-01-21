package grpc

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "github.com/win5do/golang-microservice-demo/pkg/api/petpb"

	log "github.com/win5do/go-lib/logx"
)

func runGateway(gatewayAddr, grpcAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jsonPb := &runtime.JSONPb{}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonPb),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterPetServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}

	log.Infof("gateway server start: %s", gatewayAddr)
	return http.ListenAndServe(gatewayAddr, mux)
}
