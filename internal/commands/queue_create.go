package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/queues"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var queueCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a queue",
	RunE:  createQueue,
}

func createQueue(cmd *cobra.Command, args []string) error {

	ctl, err := do.Invoke[*queues.Controller](di.Injector)
	if err != nil {
		return err
	}
	req := dto.CreateQueueRequest{
		Name:        queueName,
		DisplayName: queueName,
	}
	queue, err := ctl.CreateQueue(organizationName, projectName, req)
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
	requireOrganization(queueCreateCmd)
	requireProject(queueCreateCmd)
	requireQueue(queueCreateCmd)
	queueCmd.AddCommand(queueCreateCmd)
}
