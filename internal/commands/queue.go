package commands

import "github.com/spf13/cobra"

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manage queues",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(queueCmd)
}
