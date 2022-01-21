package main

import (
	goflag "flag"
	"github.com/spf13/cobra"

	log "github.com/win5do/go-lib/logx"
	"github.com/win5do/golang-microservice-demo/pkg/config"
	grpcserver "github.com/win5do/golang-microservice-demo/pkg/server/grpc"
	httpserver "github.com/win5do/golang-microservice-demo/pkg/server/http"
)

func main() {
	cfg := config.DefaultConfig()

	rootCmd := &cobra.Command{
		Use:   cfg.AppName,
		Short: "golang microservice demo",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := config.InitConfig(cfg)
			if err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			Run(cfg)
		},
	}

	config.SetFlags(rootCmd.Flags(), cfg)
	rootCmd.Flags().AddGoFlagSet(goflag.CommandLine)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("err: %+v", err)
	}
}

func Run(cfg *config.Config) {
	// http
	go httpserver.Run(cfg)

	// grpc
	go grpcserver.Run(cfg)

	<-make(chan struct{})
}
