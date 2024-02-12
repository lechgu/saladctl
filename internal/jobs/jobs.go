package jobs

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

func (ctl *Controller) PostJob(organization string, project string, queue string, input []byte) (dto.Job, error) {

	req := dto.CreateJobRequest{
		Input: input,
	}
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s/jobs", ctl.cfg.BaseURL, organization, project, queue)
	return sessions.CreateOne[dto.Job, dto.CreateJobRequest](ctl.session, url, req)
}

func (ctl *Controller) ListJobs(organization string, project string, queue string) ([]dto.Job, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s/jobs", ctl.cfg.BaseURL, organization, project, queue)
	return sessions.GetMany[dto.Job](ctl.session, url)
}

func (ctl *Controller) GetJob(organization string, project string, queue string, job string) (dto.Job, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s/jobs/%s", ctl.cfg.BaseURL, organization, project, queue, job)
	return sessions.GetOne[dto.Job](ctl.session, url)
}

func (ctl *Controller) DeleteJob(organization string, project string, queue string, job string) error {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s/jobs/%s", ctl.cfg.BaseURL, organization, project, queue, job)
	return sessions.DeleteOne(ctl.session, url)
}
