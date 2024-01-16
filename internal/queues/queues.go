package queues

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

func (ctl *Controller) ListQueues(organization string, project string) ([]dto.Queue, error) {
	queues := []dto.Queue{}
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues", ctl.cfg.BaseURL, organization, project)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return queues, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return queues, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return queues, err
	}
	var queueList dto.QueueList
	err = json.Unmarshal(payload, &queueList)
	if err != nil {
		return queues, err
	}
	return queueList.Items, nil
}

func (ctl *Controller) getQueue(organization string, project string, name string) (dto.Queue, error) {
	var queue dto.Queue
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s", ctl.cfg.BaseURL, organization, project, name)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return queue, err
	}
	defer res.Body.Close()
	if res.StatusCode <= 200 || res.StatusCode >= 300 {
		return queue, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return queue, err
	}
	err = json.Unmarshal(payload, &queue)
	return queue, err
}
