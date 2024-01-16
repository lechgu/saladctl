package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/organizations"
	"time"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var organizationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the Organization",
	RunE:  getOrganization,
}

func init() {
	organizationGetCmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = organizationGetCmd.MarkFlagRequired("organization")
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
	fmt.Printf("Id:                    %s\n", org.ID)
	fmt.Printf("Name:                  %s\n", org.Name)
	fmt.Printf("Display name:          %s\n", org.DisplayName)
	fmt.Printf("Create:                %s\n", org.CreateTime.Format(time.RFC822))
	fmt.Printf("Update:                %s\n", org.UpdateTime.Format(time.RFC822))
	fmt.Printf("Has had valid payment: %t\n", org.HasHadValidPayment)
	return nil
}
