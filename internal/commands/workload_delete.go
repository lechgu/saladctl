package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/workloads"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var workloadDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete workload",
	RunE:  deleteWorkload,
}

func deleteWorkload(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*workloads.Controller](di.Injector)
	if err != nil {
		return err
	}
	return ctl.DeleteWorkload(organizationName, projectName, workloadName)
}

func init() {
	requireOrganization(workloadDelete)
	requireProject(workloadDelete)
	requireWorkload(workloadDelete)
	workloadCmd.AddCommand(workloadDelete)
}
