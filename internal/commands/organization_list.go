package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/organizations"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var organizationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	RunE:  listOrganizations,
}

func listOrganizations(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*organizations.Controller](di.Injector)
	if err != nil {
		return err
	}
	orgs, err := ctl.ListOrganizations()
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Name", "Created")
	lo.ForEach(orgs, func(org dto.Organization, _ int) {
		tbl.AddRow(org.ID,
			org.Name,
			humanize.Time(org.CreateTime),
		)
	})
	tbl.Print()
	return nil
}

func init() {
	organizationCmd.AddCommand(organizationListCmd)
}
