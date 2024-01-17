package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/projects"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	RunE:  listProjects,
}

func listProjects(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*projects.Controller](di.Injector)
	if err != nil {
		return err
	}
	projects, err := ctl.ListProjects(organizationName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Name", "Created")
	lo.ForEach(projects, func(project dto.Project, _ int) {
		tbl.AddRow(project.ID,
			project.Name,
			humanize.Time(project.CreateTime),
		)
	})
	tbl.Print()
	return nil
}

func init() {
	requireOrganization(projectListCmd)
	projectCmd.AddCommand(projectListCmd)
}
