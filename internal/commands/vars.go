package commands

import "github.com/spf13/cobra"

var (
	organizationName string
	projectName      string
	workloadName     string
	queueName        string
	jobID            string
	payloadFile      string
)

func requireOrganization(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = cmd.MarkFlagRequired("organization")
}

func requireProject(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	_ = cmd.MarkFlagRequired("project")
}

func requireWorkload(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&workloadName, "workload", "w", "", "Workload name")
	_ = cmd.MarkFlagRequired("workload")
}

func requireQueue(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&queueName, "queue", "q", "", "Queue name")
	_ = cmd.MarkFlagRequired("queue")
}

func requireJob(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&jobID, "job", "j", "", "Job id")
	_ = cmd.MarkFlagRequired("job")
}

func requirePayload(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&payloadFile, "payload", "f", "", "File containing the job payload")
	_ = cmd.MarkFlagRequired("payload")
}
