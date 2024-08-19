package configuration

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type ConfigurationApiProxy struct {
	HTTPClient  *http.Client
	Auth        AuthStruct
	AccountName string
}

type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func NewProxy(username, password, account_name *string) (*ConfigurationApiProxy, error) {
	proxy := ConfigurationApiProxy{
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
		AccountName: *account_name,
		Auth: AuthStruct{
			Username: *username,
			Password: *password,
		},
	}

	response, err := proxy.GetToken(username, password)
	if err != nil {
		return nil, err
	}

	proxy.Auth.Token = response.Token

	return &proxy, nil
}

func (proxy *ConfigurationApiProxy) MakeRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("cdn-auth-token", proxy.Auth.Token)

	res, err := proxy.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
