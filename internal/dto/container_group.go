package dto

import "time"

type ContainerResources struct {
	CPU    int `json:"cpu" yaml:"cpu"`
	Memory int `json:"memory" yaml:"memory"`
}

type Container struct {
	Image                string             `json:"image"`
	Resources            ContainerResources `json:"resources" yaml:"resources"`
	Command              []string           `json:"command" yaml:"command"`
	Size                 int                `json:"size" yaml:"size"`
	EnvironmentVariables map[string]string  `json:"environment_variables"`
}

type InstanceStatusCount struct {
	AllocatingCount int `json:"allocating_count" yaml:"allocating_count"`
	CreatingCount   int `json:"creating_count" yaml:"creating_count"`
	RunningCount    int `json:"running_count" yaml:"running_count"`
}

type ContainerGroupState struct {
	Status              string              `json:"status"`
	Description         string              `json:"description" yaml:"description"`
	StartTime           time.Time           `json:"start_time"`
	FinishTime          time.Time           `json:"finish_time"`
	InstanceStatusCount InstanceStatusCount `json:"instance_status_count"`
}

type ContainerGroupQueueConnection struct {
	Path      string `json:"path" yaml:"path"`
	Port      int    `json:"port" yaml:"port"`
	QueueName string `json:"queue_name" yaml:"queue_name"`
}

type ContainerGroup struct {
	ID              string                        `json:"id" yaml:"id"`
	Name            string                        `json:"name" yaml:"name"`
	DisplayName     string                        `json:"display_name" yaml:"display_name"`
	AutostartPolicy bool                          `json:"autostart_policy" yaml:"autostart_policy"`
	Replicas        int                           `json:"replicas" yaml:"replicas"`
	RestartPolicy   string                        `json:"restart_policy" yaml:"restart_policy"`
	Container       Container                     `json:"container" yaml:"container"`
	CurrentState    ContainerGroupState           `json:"current_state" yaml:"current_state"`
	QueueConnection ContainerGroupQueueConnection `json:"queue_connection" yaml:"queue_connection"`
	CreateTime      time.Time                     `json:"create_time" yaml:"create_time"`
	UpdateTime      time.Time                     `json:"update_time" yaml:"update_time"`
	Version         int                           `json:"version" yaml:"version"`
}

type ContainerGroupList struct {
	Items []ContainerGroup `json:"items"`
}
