package organizations

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

func (ctl *Controller) ListOrganizations() ([]dto.Organization, error) {
	var orgs []dto.Organization
	url := fmt.Sprintf("%s/organizations", ctl.cfg.BaseURL)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return orgs, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return orgs, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return orgs, err
	}
	var response dto.OrganizationList
	if err = json.Unmarshal(payload, &response); err != nil {
		return orgs, err
	}
	return response.Items, err
}

func (ctl *Controller) GetOrganization(name string) (dto.Organization, error) {
	var org dto.Organization
	url := fmt.Sprintf("%s/organizations/%s", ctl.cfg.BaseURL, name)
	res, err := ctl.session.Client.Get(url)
	if err != nil {
		return org, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return org, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(payload, &org)
	return org, err
}
