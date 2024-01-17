package commands

import "github.com/spf13/cobra"

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Manage instances",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(instanceCmd)
}
