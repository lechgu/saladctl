package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/projects"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var projectGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the Project",
	RunE:  getProject,
}

func init() {
	projectGetCmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = projectCmd.MarkFlagRequired("organization")
	projectGetCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	_ = projectGetCmd.MarkFlagRequired("project")
	projectCmd.AddCommand(projectGetCmd)
}

func getProject(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*projects.Controller](di.Injector)
	if err != nil {
		return err
	}
	project, err := ctl.GetProject(organizationName, projectName)
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
