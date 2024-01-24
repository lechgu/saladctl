package workloads

import (
	"fmt"
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

func (ctl *Controller) ListWorkloads(organization string, project string) ([]dto.ContainerGroup, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers", ctl.cfg.BaseURL, organization, project)
	return sessions.GetMany[dto.ContainerGroup](ctl.session, url)
}

func (ctl *Controller) GetWorkload(organization string, project string, name string) (dto.ContainerGroup, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.GetOne[dto.ContainerGroup](ctl.session, url)
}
