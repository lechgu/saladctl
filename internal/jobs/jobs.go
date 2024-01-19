package jobs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/dto"
	"lechgu/saladctl/internal/sessions"
	"net/http"

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
	var job dto.Job
	createJobRequest := dto.CreateJobRequest{
		Input: input,
	}
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s/jobs", ctl.cfg.BaseURL, organization, project, queue)
	payload, err := json.Marshal(createJobRequest)
	if err != nil {
		return job, err
	}
	resp, err := ctl.session.Client.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return job, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return job, errors.New(resp.Status)
	}
	payload, err = io.ReadAll(resp.Body)
	if err != nil {
		return job, err
	}
	err = json.Unmarshal(payload, &job)
	return job, err
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
	req, err := http.NewRequest(http.MethodDelete, url, http.NoBody)
	if err != nil {
		return err
	}
	res, err := ctl.session.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusAccepted {
		return errors.New(res.Status)
	}
	return nil
}
