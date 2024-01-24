package commands

import "github.com/spf13/cobra"

var workloadCmd = &cobra.Command{
	Use:   "workload",
	Short: "Manage workloads",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(workloadCmd)
}
