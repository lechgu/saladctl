package sessions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lechgu/saladctl/internal/config"
	"lechgu/saladctl/internal/dto"
	"net/http"
	"net/http/cookiejar"

	"github.com/samber/do"
)

type Session struct {
	Client http.Client
	cfg    *config.Config
}

func New(di *do.Injector) (*Session, error) {
	cfg, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	sess := Session{
		cfg: cfg,
		Client: http.Client{
			Jar: jar,
		},
	}
	if err = sess.login(); err != nil {
		return nil, err
	}
	return &sess, nil
}

func (s *Session) login() error {

	url := fmt.Sprintf("%s/users/login", s.cfg.BaseURL)
	loginRequest := dto.LoginRequest{
		Email:    s.cfg.Email,
		Password: s.cfg.Password,
	}
	payload, err := json.Marshal(loginRequest)
	if err != nil {
		return err
	}
	res, err := s.Client.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New(res.Status)
	}
	defer res.Body.Close()
	return nil
}
