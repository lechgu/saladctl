package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/queues"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var queueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List queues",
	RunE:  ListQueues,
}

func ListQueues(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*queues.Controller](di.Injector)
	if err != nil {
		return err
	}
	queues, err := ctl.ListQueues(organizationName, projectName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Name", "Created")
	lo.ForEach(queues, func(queue dto.Queue, _ int) {
		tbl.AddRow(queue.ID, queue.Name,
			humanize.Time(queue.CreateTime),
		)
	})
	tbl.Print()
	return nil
}

func init() {
	queueListCmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = queueListCmd.MarkFlagRequired("organization")
	queueListCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	_ = queueListCmd.MarkFlagRequired("project")
	queueCmd.AddCommand(queueListCmd)
}
