package commands

import "github.com/spf13/cobra"

var containerGroupCmd = &cobra.Command{
	Use:   "containergroup",
	Short: "Manage container groups",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(containerGroupCmd)
}
