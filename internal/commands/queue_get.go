package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/queues"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var queueGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get queue",
	RunE:  getQueue,
}

func getQueue(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*queues.Controller](di.Injector)
	if err != nil {
		return err
	}

	queue, err := ctl.GetQueue(organizationName, projectName, queueName)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(queue)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}

func init() {
	queueGetCmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = queueGetCmd.MarkFlagRequired("organization")
	queueGetCmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	_ = queueGetCmd.MarkFlagRequired("project")
	queueGetCmd.Flags().StringVarP(&queueName, "queue", "q", "", "Queue name")
	_ = queueGetCmd.MarkFlagRequired("queue")
	queueCmd.AddCommand(queueGetCmd)
}
