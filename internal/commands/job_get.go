package commands

import (
	"fmt"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/jobs"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var jobGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the Job",
	RunE:  getJob,
}

func init() {
	requireOrganization(jobGetCmd)
	requireProject(jobGetCmd)
	requireQueue(jobGetCmd)
	requireJob(jobGetCmd)
	jobCmd.AddCommand(jobGetCmd)
}

func getJob(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*jobs.Controller](di.Injector)
	if err != nil {
		return err
	}
	job, err := ctl.GetJob(organizationName, projectName, queueName, jobID)
	if err != nil {
		return err
	}
	dump, err := yaml.Marshal(job.MakePretty())
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}
