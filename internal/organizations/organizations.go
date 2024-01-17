package organizations

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

func (ctl *Controller) ListOrganizations() ([]dto.Organization, error) {
	url := fmt.Sprintf("%s/organizations", ctl.cfg.BaseURL)
	return sessions.GetMany[dto.Organization](ctl.session, url)
}

func (ctl *Controller) GetOrganization(name string) (dto.Organization, error) {

	url := fmt.Sprintf("%s/organizations/%s", ctl.cfg.BaseURL, name)
	return sessions.GetOne[dto.Organization](ctl.session, url)
}
