package commands

import (
	"fmt"
	"lechgu/saladctl/internal/containergroups"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var containerGroupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List container groups",
	RunE:  listContainerGroups,
}

func listContainerGroups(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*containergroups.Controller](di.Injector)
	if err != nil {
		return err
	}
	containerGroups, err := ctl.ListContainerGroups(organizationName, projectName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Name", "Image", "Status", "Instances", "Created")
	lo.ForEach(containerGroups, func(containerGroup dto.ContainerGroup, _ int) {
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
	containerGroupListCmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	containerGroupListCmd.MarkFlagRequired("organization")
	containerGroupListCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	containerGroupListCmd.MarkFlagRequired("project")
	containerGroupCmd.AddCommand(containerGroupListCmd)
}
