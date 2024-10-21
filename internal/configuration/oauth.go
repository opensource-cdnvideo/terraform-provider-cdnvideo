package configuration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const OauthURL string = "https://api.cdnvideo.ru/app/oauth/v1/token/"

type AuthResponse struct {
	Status   int    `json:"status"`
	Lifetime int    `json:"lifetime"`
	Token    string `json:"token"`
}

func (proxy *ConfigurationApiProxy) GetToken(username, password *string) (*AuthResponse, error) {
	data := url.Values{}
	data.Set("username", *username)
	data.Set("password", *password)
	req, err := http.NewRequest("POST", OauthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	body, err := proxy.MakeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	var ar AuthResponse
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if ar.Status != http.StatusOK {
		return nil, fmt.Errorf("authentication failed with status: %d", ar.Status)
	}

	return &ar, nil
}
