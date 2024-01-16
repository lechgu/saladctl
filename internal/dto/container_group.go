package dto

import "time"

type Container struct {
	Image                string            `json:"image"`
	EnvironmentVariables map[string]string `json:"environment_variables"`
}

type InstanceStatusCount struct {
	AllocatingCount int `json:"allocating_count"`
	CreatingCount   int `json:"creating_count"`
	RunningCount    int `json:"running_count"`
}

type ContainerGroupState struct {
	Status              string              `json:"status"`
	StartTime           time.Time           `json:"start_time"`
	FinishTime          time.Time           `json:"finish_time"`
	InstanceStatusCount InstanceStatusCount `json:"instance_status_count"`
}

type ContainerGroup struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	DisplayName     string              `json:"display_name"`
	CreateTime      time.Time           `json:"create_time"`
	UpdateTime      time.Time           `json:"update_time"`
	AutostartPolicy bool                `json:"autostart_policy"`
	RestartPolicy   string              `json:"restart_policy"`
	Replicas        int                 `json:"replicas"`
	Version         int                 `json:"version"`
	Container       Container           `json:"container"`
	CurrentState    ContainerGroupState `json:"current_state"`
}

type ContainerGroupList struct {
	Items []ContainerGroup `json:"items"`
}
