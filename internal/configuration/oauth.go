package configuration

import (
	"encoding/json"
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
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
