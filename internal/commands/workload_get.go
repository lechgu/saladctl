package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/workloads"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var workloadGet = &cobra.Command{
	Use:   "get",
	Short: "Get workload",
	RunE:  getWorkload,
}

func getWorkload(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*workloads.Controller](di.Injector)
	if err != nil {
		return err
	}
	workload, err := ctl.GetWorkload(organizationName, projectName, workloadName)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(workload)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}

func init() {
	requireOrganization(workloadGet)
	requireProject(workloadGet)
	requireWorkload(workloadGet)
	workloadCmd.AddCommand(workloadGet)
}
