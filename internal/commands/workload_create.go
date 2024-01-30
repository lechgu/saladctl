package commands

import (
	"github.com/spf13/cobra"
)

var workloadCreate = &cobra.Command{
	Use:   "create",
	Short: "Create workload",
	RunE:  createWorkload,
}

func createWorkload(cmd *cobra.Command, args []string) error {
	_ = cmd.Help()
	return nil
}

func init() {
	requireOrganization(workloadCreate)
	requireProject(workloadCreate)
	workloadCmd.AddCommand(workloadCreate)
}
