package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/instances"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var instanceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List instances",
	RunE:  ListInstances,
}

func ListInstances(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*instances.Controller](di.Injector)
	if err != nil {
		return err
	}
	instances, err := ctl.ListInstances(organizationName, projectName, containerGroupName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "State", "Updated")
	lo.ForEach(instances, func(inst dto.Instance, _ int) {
		tbl.AddRow(inst.MachineID,
			inst.State,
			humanize.Time(inst.UpdateTime),
		)
	})
	tbl.Print()
	return nil
}

func init() {
	requireOrganization(instanceListCmd)
	requireProject(instanceListCmd)
	requireContainerGroup(instanceListCmd)
	instanceCmd.AddCommand(instanceListCmd)
}
