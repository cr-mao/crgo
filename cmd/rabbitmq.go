package cmd

import "github.com/spf13/cobra"
import "crgo/infra/rabbitmq"

func init() {
	RabbitMqTestCmd.AddCommand(PublishTestCmd)
	RabbitMqTestCmd.AddCommand(ConsumeCmd)
}

var RabbitMqTestCmd = &cobra.Command{
	Use:   "rabbitmq",
	Short: "rabbitmq生产者",
	Long:  "rabbitmq生产者",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var PublishTestCmd = &cobra.Command{
	Use:   "publish",
	Short: "rabbitmq生产者",
	Long:  "rabbitmq生产者",
	Run: func(cmd *cobra.Command, args []string) {
		rabbitmq.Pub()
	},
}

var ConsumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "rabbitmq消费者",
	Long:  "rabbitmq消费者",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
