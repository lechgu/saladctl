package instances

import (
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

func (ctl *Controller) ListInstances(organization string, project string, containerGroup string) ([]dto.Instance, error) {
	url := fmt.Sprintf("%s/organizations/%s/projects/%s/containers/%s/instances", ctl.cfg.BaseURL, organization, project, containerGroup)
	var items []dto.Instance
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return items, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return items, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return items, err
	}
	var coll dto.InstanceList
	err = json.Unmarshal(payload, &coll)
	return coll.Instances, err
}
