package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/projects"
	"time"

	"github.com/samber/do"
	"github.com/spf13/cobra"
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
	fmt.Printf("Id:                    %s\n", project.ID)
	fmt.Printf("Name:                  %s\n", project.Name)
	fmt.Printf("Display name:          %s\n", project.DisplayName)
	fmt.Printf("Create:                %s\n", project.CreateTime.Format(time.RFC822))
	fmt.Printf("Update:                %s\n", project.UpdateTime.Format(time.RFC822))
	fmt.Printf("Has had valid payment: %t\n", project.HasHadValidPayment)
	return nil
}
