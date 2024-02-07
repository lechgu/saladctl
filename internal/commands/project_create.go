package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/projects"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	RunE:  createProject,
}

func createProject(cmd *cobra.Command, args []string) error {

	ctl, err := do.Invoke[*projects.Controller](di.Injector)
	if err != nil {
		return err
	}
	req := dto.CreateProjectRequest{
		Name: projectName,
	}
	project, err := ctl.CreateProject(organizationName, req)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(project)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}

func init() {
	requireOrganization(projectCreateCmd)
	requireProject(projectCreateCmd)
	projectCmd.AddCommand(projectCreateCmd)
}
