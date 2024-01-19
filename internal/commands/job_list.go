package commands

import (
	"lechgu/saladctl/internal/di"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/jobs"

	"github.com/dustin/go-humanize"
	"github.com/rodaine/table"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var jobListCmd = &cobra.Command{
	Use:   "list",
	Short: "List jobs",
	RunE:  listJobs,
}

func init() {
	requireOrganization(jobListCmd)
	requireProject(jobListCmd)
	requireQueue(jobListCmd)
	jobCmd.AddCommand(jobListCmd)
}

func listJobs(cmd *cobra.Command, args []string) error {
	ctl, err := do.Invoke[*jobs.Controller](di.Injector)
	if err != nil {
		return err
	}
	jobs, err := ctl.ListJobs(organizationName, projectName, queueName)
	if err != nil {
		return err
	}
	tbl := table.New("Id", "Status", "Created")
	lo.ForEach(jobs, func(job dto.Job, _ int) {
		tbl.AddRow(job.ID,
			job.Status,
			humanize.Time(job.CreateTime),
		)
	})
	tbl.Print()
	return nil
}
