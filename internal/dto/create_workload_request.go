package dto

type CreateWorkloadRequest struct {
	Name                 string                  `json:"name"`
	Container            Container               `json:"container"`
	Command              []string                `json:"command"`
	EnvironmentVariables map[string]string       `json:"environment_variables"`
	QueueConnection      WorkloadQueueConnection `json:"queue_connection"`
}
