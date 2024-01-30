package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/queues"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var queueDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete queue",
	RunE:  deleteQueue,
}

func deleteQueue(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*queues.Controller](di.Injector)
	if err != nil {
		return err
	}

	return ctl.DeleteQueue(organizationName, projectName, queueName)
}

func init() {
	requireOrganization(queueDeleteCmd)
	requireProject(queueDeleteCmd)
	requireQueue(queueDeleteCmd)
	queueCmd.AddCommand(queueDeleteCmd)
}
