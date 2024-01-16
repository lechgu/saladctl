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
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New(resp.Status)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	return nil
}
