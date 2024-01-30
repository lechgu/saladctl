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

func (ctl *Controller) ListWorkloads(organization string, project string) ([]dto.Workload, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers", ctl.cfg.BaseURL, organization, project)
	return sessions.GetMany[dto.Workload](ctl.session, url)
}

func (ctl *Controller) GetWorkload(organization string, project string, name string) (dto.Workload, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.GetOne[dto.Workload](ctl.session, url)
}

func (ctl *Controller) DeleteWorkload(organization string, project string, name string) error {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.DeleteOne(ctl.session, url)
}
