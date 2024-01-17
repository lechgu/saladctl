package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/organizations"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var organizationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the Organization",
	RunE:  getOrganization,
}

func init() {
	requireOrganization(organizationGetCmd)
	organizationCmd.AddCommand(organizationGetCmd)
}

func getOrganization(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*organizations.Controller](di.Injector)
	if err != nil {
		return err
	}
	org, err := ctl.GetOrganization(organizationName)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(org)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}
