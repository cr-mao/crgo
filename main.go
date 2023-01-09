package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"crgo/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "crgo",
	Short: "This is crgo",
}

func init() {
	//注入版本号，版本时间
	rootCmd.AddCommand(cmd.WordCmd)
	rootCmd.AddCommand(cmd.TimeCmd)
	rootCmd.AddCommand(cmd.SqlCmd)
	rootCmd.AddCommand(cmd.ServeCmd)
	rootCmd.AddCommand(cmd.GrpcClientCmd)
	rootCmd.AddCommand(cmd.RabbitMqTestCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
