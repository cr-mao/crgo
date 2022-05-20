package cmd

import (
	"github.com/spf13/cobra"

	"crgo/grpc/client"
	"crgo/infra/log"
)

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
