package containergroups

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/sessions"

	"github.com/samber/do"
)

type Controller struct {
	cfg     *config.Config
	session *sessions.Session
}

func NewController(di *do.Injector) (*Controller, error) {
	cfg, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	session, err := do.Invoke[*sessions.Session](di)
	if err != nil {
		return nil, err
	}
	return &Controller{
		cfg:     cfg,
		session: session,
	}, nil
}

func (ctl *Controller) ListContainerGroups(organization string, project string) ([]dto.ContainerGroup, error) {
	containerGroups := []dto.ContainerGroup{}
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers", ctl.cfg.BaseURL, organization, project)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return containerGroups, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return containerGroups, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return containerGroups, err
	}
	var containerGroupsList dto.ContainerGroupList
	if err = json.Unmarshal(payload, &containerGroupsList); err != nil {
		return containerGroups, err
	}
	return containerGroupsList.Items, nil
}

func (ctl *Controller) GetContainerGroup(organization string, project string, name string) (dto.ContainerGroup, error) {
	var containerGroup dto.ContainerGroup
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers/%s", ctl.cfg.BaseURL, organization, project, name)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return containerGroup, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return containerGroup, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return containerGroup, err
	}
	err = json.Unmarshal(payload, &containerGroup)
	return containerGroup, err
}
