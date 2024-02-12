package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workloadCreate = &cobra.Command{
	Use:   "create",
	Short: "Create workload",
	RunE:  createWorkload,
}

func createWorkload(cmd *cobra.Command, args []string) error {
	fmt.Println("creating workflow...")
	return nil
}

func init() {
	requireOrganization(workloadCreate)
	requireProject(workloadCreate)
	requireWorkload(workloadCreate)
	requireImage(workloadCreate)
	workloadCmd.AddCommand(workloadCreate)
}
