package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/jobs"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var jobResultCmd = &cobra.Command{
	Use:   "result",
	Short: "Get the job result",
	RunE:  getResult,
}

func init() {
	requireOrganization(jobResultCmd)
	requireProject(jobResultCmd)
	requireQueue(jobResultCmd)
	requireJob(jobResultCmd)
	jobCmd.AddCommand(jobResultCmd)
}

func getResult(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*jobs.Controller](di.Injector)
	if err != nil {
		return err
	}
	job, err := ctl.GetJob(organizationName, projectName, queueName, jobID)
	if err != nil {
		return err
	}
	fmt.Println(string(job.Output))
	return nil
}
