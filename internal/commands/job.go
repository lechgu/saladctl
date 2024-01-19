package commands

import "github.com/spf13/cobra"

var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Manage jobs",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
}
