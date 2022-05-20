package cmd

import (
	"github.com/spf13/cobra"

	"crgo/httpproject"
	"crgo/infra/log"
)

var HttpServeCmd = &cobra.Command{
	Use:   "httpserve",
	Short: "http serve",
	Long:  desc,
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := httpproject.Run(); err != nil {
			log.Error(err)
		}
	},
}
