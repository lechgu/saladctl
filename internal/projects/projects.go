package projects

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

func (ctl *Controller) ListProjects(organization string) ([]dto.Project, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects", ctl.cfg.BaseURL, organization)
	return sessions.GetMany[dto.Project](ctl.session, url)
}

func (ctl *Controller) GetProject(organization string, name string) (dto.Project, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s", ctl.cfg.BaseURL, organization, name)
	return sessions.GetOne[dto.Project](ctl.session, url)
}

func (ctl *Controller) CreateProject(organization string, req dto.CreateProjectRequest) (dto.Project, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects", ctl.cfg.BaseURL, organization)
	return sessions.CreateOne[dto.Project, dto.CreateProjectRequest](ctl.session, url, req)
}

func (ctl *Controller) DeleteProject(organization string, project string) error {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s", ctl.cfg.BaseURL, organization, project)
	return sessions.DeleteOne(ctl.session, url)
}
