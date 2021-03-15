package cmd

import (
	"github.com/spf13/cobra"

	"crgo/grpcproject"
	"crgo/grpcproject/client"
	"crgo/infra/log"
)

var GrpcServeCmd = &cobra.Command{
	Use:   "grpcserve",
	Short: "grpc serve",
	Long:  desc,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := grpcproject.Run(); err != nil {
			log.Error(err)
		}
	},
}

var GrpcClientCmd = &cobra.Command{
	Use:   "grpcclient",
	Short: "grpc client",
	Long:  desc,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := client.Do(); err != nil {
			log.Error(err)
		}
	},
}
