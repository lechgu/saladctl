package commands

import "github.com/spf13/cobra"

var (
	organizationName   string
	projectName        string
	containerGroupName string
	queueName          string
	jobID              string
	path               string
	project            int
)

func requireOrganization(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&organizationName, "organization", "o", "", "Organization name")
	_ = cmd.MarkFlagRequired("organization")
}

func requireProject(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&projectName, "project", "p", "", "Project name")
	_ = cmd.MarkFlagRequired("project")
}

func requireContainerGroup(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&containerGroupName, "container-group", "c", "", "Container group name")
	_ = cmd.MarkFlagRequired("container-group")
}

func requireQueue(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&queueName, "queue", "q", "", "Queue name")
	_ = cmd.MarkFlagRequired("queue")
}

func requireJob(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&jobID, "job", "j", "", "Job id")
	_ = cmd.MarkFlagRequired("job")
}
