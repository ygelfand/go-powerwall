package powerwall

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (p *PowerwallGateway) makeAPIRequest(method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, p.endpoint.JoinPath("api", path).String(), body)
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	err = retry(5, 15*time.Millisecond, func() error {

		if p.authToken != "" {
			req.AddCookie(&http.Cookie{
				Name:  "AuthCookie",
				Value: p.authToken,
			})
		} else {
			err = p.refreshAuthToken()
			if err != nil {
				// auth failed
				return stop{err}
			}
			return errors.New("retrying with auth")
		}

		req.Header.Set("Content-type", "application/json")
		resp, err = p.httpClient.Do(req)
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			err = p.refreshAuthToken()
			if err != nil {
				// auth failed
				return stop{err}
			}
			return errors.New("retrying with auth")

		}
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respbody, nil
}

func (p *PowerwallGateway) refreshAuthToken() error {
	fmt.Println("Refreshing auth token")
	auth := map[string]string{"username": "customer", "password": p.password[len(p.password)-5:]}
	jsonAuth, _ := json.Marshal(auth)
	req, err := http.NewRequest("POST", p.endpoint.JoinPath("api/login/Basic").String(), bytes.NewBuffer(jsonAuth))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	loginResp := &loginResponse{}
	err = json.Unmarshal(respbody, loginResp)
	if err != nil {
		return err
	}
	p.authToken = loginResp.Token
	return nil
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}
