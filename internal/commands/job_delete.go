package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/jobs"

	"github.com/samber/do"
	"github.com/spf13/cobra"
)

var jobDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the Job",
	RunE:  deleteJob,
}

func init() {
	requireOrganization(jobDeleteCmd)
	requireProject(jobDeleteCmd)
	requireQueue(jobDeleteCmd)
	requireJob(jobDeleteCmd)
	jobCmd.AddCommand(jobDeleteCmd)
}

func deleteJob(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*jobs.Controller](di.Injector)
	if err != nil {
		return err
	}
	return ctl.DeleteJob(organizationName, projectName, queueName, jobID)
}
