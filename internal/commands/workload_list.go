package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/workloads"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var workloadList = &cobra.Command{
	Use:   "list",
	Short: "List workloads",
	RunE:  listWorkloads,
}

func listWorkloads(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*workloads.Controller](di.Injector)
	if err != nil {
		return err
	}
	workloads, err := ctl.ListWorkloads(organizationName, projectName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Name", "Image", "Status", "Instances", "Created")
	lo.ForEach(workloads, func(containerGroup dto.Workload, _ int) {
		tbl.AddRow(containerGroup.ID,
			containerGroup.Name,
			containerGroup.Container.Image,
			containerGroup.CurrentState.Status,
			fmt.Sprintf("%d/%d/%d", containerGroup.CurrentState.InstanceStatusCount.AllocatingCount, containerGroup.CurrentState.InstanceStatusCount.CreatingCount, containerGroup.CurrentState.InstanceStatusCount.RunningCount),
			humanize.Time(containerGroup.CreateTime),
		)
	})
	tbl.Print()
	return nil
}

func init() {
	requireOrganization(workloadList)
	requireProject(workloadList)
	workloadCmd.AddCommand(workloadList)
}
