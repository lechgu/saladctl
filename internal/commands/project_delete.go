package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/projects"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var projectDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the Project",
	RunE:  deleteProject,
}

func init() {
	requireOrganization(projectDeleteCmd)
	requireProject(projectDeleteCmd)
	projectCmd.AddCommand(projectDeleteCmd)
}

func deleteProject(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*projects.Controller](di.Injector)
	if err != nil {
		return err
	}
	return ctl.DeleteProject(organizationName, projectName)
}
