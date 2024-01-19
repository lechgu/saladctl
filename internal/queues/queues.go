package queues

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

func (ctl *Controller) ListQueues(organization string, project string) ([]dto.Queue, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues", ctl.cfg.BaseURL, organization, project)
	return sessions.GetMany[dto.Queue](ctl.session, url)
}

func (ctl *Controller) GetQueue(organization string, project string, name string) (dto.Queue, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues/%s", ctl.cfg.BaseURL, organization, project, name)
	return sessions.GetOne[dto.Queue](ctl.session, url)
}

func (ctl *Controller) CreateQueue(organization string, project string, req dto.QueueCreateRequest) (dto.Queue, error) {
	var queue dto.Queue
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/queues", ctl.cfg.BaseURL, organization, project)
	payload, err := json.Marshal(req)
	if err != nil {
		return queue, err
	}
	res, err := ctl.session.Client.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return queue, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return queue, errors.New(res.Status)
	}
	payload, err = io.ReadAll(res.Body)
	if err != nil {
		return queue, err
	}
	err = json.Unmarshal(payload, &queue)
	return queue, err
}
