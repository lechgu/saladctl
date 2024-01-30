package sessions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func GetOne[T any](session *Session, url string) (T, error) {
	var instance T
	res, err := session.Client.Get(url)
	if err != nil {
		return instance, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return instance, errors.New(res.Status)
	}
	payload, err := io.ReadAll(res.Body)
	if err != nil {
		return instance, err
	}
	err = json.Unmarshal(payload, &instance)
	return instance, err
}

type Collection[T any] struct {
	Items []T `json:"items"`
}

func GetMany[T any](session *Session, url string) ([]T, error) {
	var items []T
	res, err := session.Client.Get(url)
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
	var coll Collection[T]
	err = json.Unmarshal(payload, &coll)
	return coll.Items, err
}

func DeleteOne(session *Session, url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, http.NoBody)
	if err != nil {
		return err
	}
	res, err := session.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusAccepted {
		return errors.New(res.Status)
	}
	return nil
}
