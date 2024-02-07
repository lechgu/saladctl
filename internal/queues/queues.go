package queues

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

func (ctl *Controller) ListQueues(organization string, project string) ([]dto.Queue, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues", ctl.cfg.BaseURL, organization, project)
	return sessions.GetMany[dto.Queue](ctl.session, url)
}

func (ctl *Controller) GetQueue(organization string, project string, name string) (dto.Queue, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.GetOne[dto.Queue](ctl.session, url)
}

func (ctl *Controller) DeleteQueue(organization string, project string, name string) error {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.DeleteOne(ctl.session, url)
}

func (ctl *Controller) CreateQueue(organization string, project string, req dto.CreateQueueRequest) (dto.Queue, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues", ctl.cfg.BaseURL, organization, project)
	return sessions.CreateOne[dto.Queue, dto.CreateQueueRequest](ctl.session, url, req)
}
