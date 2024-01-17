package commands

import (
	"fmt"
	"lechgu/saladctl/internal/containergroups"
	"lechgu/saladctl/internal/di"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var containerGroupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get container group",
	RunE:  getContainerGroup,
}

func getContainerGroup(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*containergroups.Controller](di.Injector)
	if err != nil {
		return err
	}
	containerGroup, err := ctl.GetContainerGroup(organizationName, projectName, containerGroupName)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(containerGroup)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}

func init() {
	requireOrganization(containerGroupGetCmd)
	requireProject(containerGroupGetCmd)
	requireContainerGroup(containerGroupCmd)
}
