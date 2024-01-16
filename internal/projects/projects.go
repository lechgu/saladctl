package projects

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

func (ctl *Controller) ListProjects(organization string) ([]dto.Project, error) {
	var projects []dto.Project
	url := fmt.Sprintf("%s/organizations/%s/projects", ctl.cfg.BaseURL, organization)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return projects, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return projects, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return projects, err
	}
	var response dto.ProjectList
	if err = json.Unmarshal(payload, &response); err != nil {
		return projects, err
	}
	return response.Items, err
}

func (ctl *Controller) GetProject(organization string, name string) (dto.Project, error) {
	var project dto.Project
	url := fmt.Sprintf("%s/organizations/%s/projects/%s", ctl.cfg.BaseURL, organization, name)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return project, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return project, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return project, err
	}
	err = json.Unmarshal(payload, &project)
	return project, err
}
