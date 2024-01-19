package commands

import (
	"fmt"
	"io"
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/jobs"
	"os"

	"github.com/samber/do"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var jobCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a job",
	RunE:  createJob,
}

func createJob(cmd *cobra.Command, args []string) error {
	payload, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	ctl, err := do.Invoke[*jobs.Controller](di.Injector)
	if err != nil {
		return err
	}
	job, err := ctl.PostJob(organizationName, projectName, queueName, payload)
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

func init() {
	requireOrganization(jobCreateCmd)
	requireProject(jobCreateCmd)
	requireQueue(jobCreateCmd)
	jobCmd.AddCommand(jobCreateCmd)
}
