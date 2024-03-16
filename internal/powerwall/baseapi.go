package powerwall

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func (p *PowerwallGateway) MakeAPIRequest(method, path string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, p.Endpoint.JoinPath("api", path).String(), body)
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

func (p *PowerwallGateway) GetAuthHeaders() []*http.Cookie {
	if p.authToken == "" {
		p.refreshAuthToken()
	}
	return []*http.Cookie{
		&http.Cookie{
			Name:  "AuthCookie",
			Value: p.authToken,
			Path:  "/",
		},
		&http.Cookie{
			Name:  "UserRecord",
			Value: p.userRecord,
			Path:  "/",
		}}
}
func (p *PowerwallGateway) refreshAuthToken() error {
	if !p.authSem.TryAcquire(1) {
		log.Println("auth refresh skipped")
		return errors.New("auth already in progress")
	}
	defer p.authSem.Release(1)
	log.Println("Refreshing auth token")
	auth := map[string]string{"username": "customer", "email": "foo@example.test", "password": p.password[len(p.password)-5:]}
	jsonAuth, _ := json.Marshal(auth)
	req, err := http.NewRequest("POST", p.Endpoint.JoinPath("api/login/Basic").String(), bytes.NewBuffer(jsonAuth))
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
	if resp.StatusCode == 429 {
		return stop{errors.New("Api throttled")}
	}
	loginResp := &loginResponse{}
	err = json.Unmarshal(respbody, loginResp)
	if err != nil {
		return err
	}
	p.authToken = loginResp.Token
	for _, c := range resp.Cookies() {
		if c.Name == "UserRecord" {
			p.userRecord = c.Value
		}
	}
	return nil
}
