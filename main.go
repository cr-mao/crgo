package main

import (
	"crgo/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crgo",
	Short: "This is crgo",
}

func init() {
	rootCmd.AddCommand(cmd.WordCmd)
	rootCmd.AddCommand(cmd.TimeCmd)
	rootCmd.AddCommand(cmd.SqlCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
