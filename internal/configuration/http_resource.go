package configuration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const ConfigurationApiURL string = "https://api.cdnvideo.ru/cdn/api/v1/%s/resource/http/%s"

type CdnHttpResource struct {
	ID                 string               `json:"id,omitempty"`
	Name               string               `json:"name,omitempty"`
	CreationTs         int64                `json:"creation_ts,omitempty"`
	CdnDomain          string               `json:"cdn_domain,omitempty"`
	Active             *bool                `json:"active,omitempty"`
	Origin             *Origin              `json:"origin,omitempty"`
	Cache              *Cache               `json:"cache,omitempty"`
	Certificate        *int64               `json:"certificate,omitempty"`
	Tuning             *string              `json:"tuning,omitempty"`
	SliceSizeMegabytes *int64               `json:"slice_size_megabytes,omitempty"`
	ModernTlsOnly      *bool                `json:"modern_tls_only,omitempty"`
	StrongSslCiphers   *bool                `json:"strong_ssl_ciphers,omitempty"`
	FollowRedirects    *bool                `json:"follow_redirects,omitempty"`
	NoHttp2            *bool                `json:"no_http2,omitempty"`
	Http2Https         *bool                `json:"http2https,omitempty"`
	HttpsOnly          *bool                `json:"https_only,omitempty"`
	UseHttp3           *bool                `json:"use_http3,omitempty"`
	Compress           *Compress            `json:"compress,omitempty"`
	Robots             *Robots              `json:"robots,omitempty"`
	Auth               *Auth                `json:"auth,omitempty"`
	Headers            *Headers             `json:"headers,omitempty"`
	Cors               *Cors                `json:"cors,omitempty"`
	Names              []string             `json:"names,omitempty"`
	Limitations        *Limitations         `json:"limitations,omitempty"`
	IOSS               *bool                `json:"ioss,omitempty"`
	Packaging          *Packaging           `json:"packaging,omitempty"`
	Locations          map[string]Locations `json:"locations,omitempty"`
}

type CdnHttpResourceCreated struct {
	Status      string `json:"status"`
	TaskId      string `json:"task_id"`
	ResourceId  string `json:"resource_id"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

type Origin struct {
	Servers        map[string]Servers `json:"servers,omitempty" tfsdk:"servers"`
	Hostname       *string            `json:"hostname,omitempty" tfsdk:"hostname"`
	HTTPS          *bool              `json:"https,omitempty" tfsdk:"https"`
	SNIHostname    *string            `json:"sni_hostname,omitempty" tfsdk:"sni_hostname"`
	ReadTimeout    *string            `json:"read_timeout,omitempty" tfsdk:"read_timeout"`
	SendTimeout    *string            `json:"send_timeout,omitempty" tfsdk:"send_timeout"`
	ConnectTimeout *string            `json:"connect_timeout,omitempty" tfsdk:"connect_timeout"`
	AWS            *AWS               `json:"aws,omitempty" tfsdk:"aws"`
	S3Bucket       *string            `json:"s3_bucket,omitempty" tfsdk:"s3_bucket"`
	SSLVerify      *bool              `json:"ssl_verify,omitempty" tfsdk:"ssl_verify"`
}

type AWS struct {
	Auth *struct {
		AccessKey *string `json:"access_key,omitempty" tfsdk:"access_key"`
		SecretKey *string `json:"secret_key,omitempty" tfsdk:"secret_key"`
	} `json:"auth,omitempty" tfsdk:"auth"`
}

type Servers struct {
	Port     *int  `json:"port,omitempty" tfsdk:"port"`
	Weight   *int  `json:"weight,omitempty" tfsdk:"weight"`
	MaxFails *int  `json:"max_fails,omitempty" tfsdk:"max_fails"`
	Backup   *bool `json:"backup,omitempty" tfsdk:"backup"`
}

type Cache struct {
	Disable          *bool     `json:"disable,omitempty" tfsdk:"disable"`
	ConsiderArgs     *bool     `json:"consider_args,omitempty" tfsdk:"consider_args"`
	ArgsWhitelist    *[]string `json:"args_whitelist,omitempty" tfsdk:"args_whitelist"`
	ConsiderCookies  *bool     `json:"consider_cookies,omitempty" tfsdk:"consider_cookies"`
	CookiesWhitelist *[]string `json:"cookies_whitelist,omitempty" tfsdk:"cookies_whitelist"`
	Valid            *struct {
		C2xx  *string `json:"2xx,omitempty" tfsdk:"c_2xx"`
		C3xx  *string `json:"3xx,omitempty" tfsdk:"c_3xx"`
		C4xx  *string `json:"4xx,omitempty" tfsdk:"c_4xx"`
		C5xx  *string `json:"5xx,omitempty" tfsdk:"c_5xx"`
		Force *bool   `json:"force,omitempty" tfsdk:"force"`
	} `json:"valid,omitempty" tfsdk:"valid"`
	UseStale *bool `json:"use_stale,omitempty" tfsdk:"use_stale"`
}

type Compress struct {
	Brotli *bool `json:"brotli,omitempty" tfsdk:"brotli"`
	Gzip   *bool `json:"gzip,omitempty" tfsdk:"gzip"`
}

type Robots struct {
	Type          *string `json:"type,omitempty" tfsdk:"type"`
	RobotsContent *string `json:"robotsContent,omitempty" tfsdk:"robots_content"`
}

type Auth struct {
	URL       *string `json:"url,omitempty" tfsdk:"url"`
	Forbidden *bool   `json:"forbidden,omitempty" tfsdk:"forbidden"`
	Md5       *struct {
		Secret   *string `json:"secret,omitempty" tfsdk:"secret"`
		Forever  *bool   `json:"forever,omitempty" tfsdk:"forever"`
		Anywhere *bool   `json:"anywhere,omitempty" tfsdk:"anywhere"`
	} `json:"md5,omitempty" tfsdk:"md5"`
}

type Headers struct {
	Request        map[string]string `json:"request,omitempty" tfsdk:"request"`
	Response       map[string]string `json:"response,omitempty" tfsdk:"response"`
	HideInResponse *[]string         `json:"hide_in_response,omitempty" tfsdk:"hide_in_response"`
}

type Cors struct {
	Domains     *[]string `json:"domains,omitempty" tfsdk:"domains"`
	Headers     *[]string `json:"headers,omitempty" tfsdk:"headers"`
	Expose      *[]string `json:"expose,omitempty" tfsdk:"expose"`
	Methods     *[]string `json:"methods,omitempty" tfsdk:"methods"`
	Credentials *bool     `json:"credentials,omitempty" tfsdk:"credentials"`
	MaxAge      *int64    `json:"max_age,omitempty" tfsdk:"max_age"`
	Disable     *bool     `json:"disable,omitempty" tfsdk:"disable"`
}

type Times struct {
	Start *string `json:"start,omitempty" tfsdk:"start"`
	End   *string `json:"end,omitempty" tfsdk:"end"`
}

type GeoLimitations struct {
	Exclude *[]struct {
		Action  *string `json:"action,omitempty" tfsdk:"action"`
		Country *string `json:"country,omitempty" tfsdk:"country"`
		Region  *string `json:"region,omitempty" tfsdk:"region"`
	} `json:"exclude,omitempty" tfsdk:"exclude"`
	DefaultAction *string  `json:"default_action,omitempty" tfsdk:"default_action"`
	Times         *[]Times `json:"times,omitempty" tfsdk:"times"`
}

type IPLimitations struct {
	Exclude *[]struct {
		IP *string `json:"ip,omitempty" tfsdk:"ip"`
	} `json:"exclude,omitempty" tfsdk:"exclude"`
	DefaultAction *string  `json:"default_action,omitempty" tfsdk:"default_action"`
	Times         *[]Times `json:"times,omitempty" tfsdk:"times"`
}

type RefererLimitations struct {
	Exclude *[]struct {
		Referer *string `json:"referer,omitempty" tfsdk:"referer"`
	} `json:"exclude,omitempty" tfsdk:"exclude"`
	DefaultAction *string  `json:"default_action,omitempty" tfsdk:"default_action"`
	Times         *[]Times `json:"times,omitempty" tfsdk:"times"`
}

type UserAgentLimitations struct {
	Exclude *[]struct {
		UserAgent *string `json:"useragent,omitempty" tfsdk:"useragent"`
	} `json:"exclude,omitempty" tfsdk:"exclude"`
	DefaultAction *string  `json:"default_action,omitempty" tfsdk:"default_action"`
	Times         *[]Times `json:"times,omitempty" tfsdk:"times"`
}

type Limitations struct {
	Geo       *[]GeoLimitations       `json:"geo,omitempty" tfsdk:"geo"`
	IP        *[]IPLimitations        `json:"ip,omitempty" tfsdk:"ip"`
	Referer   *[]RefererLimitations   `json:"referer,omitempty" tfsdk:"referer"`
	UserAgent *[]UserAgentLimitations `json:"useragent,omitempty" tfsdk:"useragent"`
}

type Locations struct {
	Cache                *Cache       `json:"cache,omitempty" tfsdk:"cache"`
	Origin               *Origin      `json:"origin,omitempty" tfsdk:"origin"`
	Auth                 *Auth        `json:"auth,omitempty" tfsdk:"auth"`
	Headers              *Headers     `json:"headers,omitempty" tfsdk:"headers"`
	Cors                 *Cors        `json:"cors,omitempty" tfsdk:"cors"`
	Limitations          *Limitations `json:"limitations,omitempty" tfsdk:"limitations"`
	IOSS                 *bool        `json:"ioss,omitempty" tfsdk:"ioss"`
	Packaging            *Packaging   `json:"packaging,omitempty" tfsdk:"packaging"`
	Rewrite              *[]Rewrite   `json:"rewrite,omitempty" tfsdk:"rewrite"`
	Compress             *Compress    `json:"compress,omitempty" tfsdk:"compress"`
	ReturnHTTPStatusCode *int         `json:"return_http_status_code,omitempty" tfsdk:"return_http_status_code"`
}

type Packaging struct {
	Mp4 *struct {
		OutputProtocols *[]string `json:"output_protocols,omitempty" tfsdk:"output_protocols"`
	} `json:"mp4,omitempty" tfsdk:"mp4"`
}
type Rewrite struct {
	From *string `json:"from,omitempty" tfsdk:"from"`
	To   *string `json:"to,omitempty" tfsdk:"to"`
	Flag *string `json:"flag,omitempty" tfsdk:"flag"`
}

func (proxy *ConfigurationApiProxy) GetHttpResources() ([]CdnHttpResource, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(ConfigurationApiURL, proxy.AccountName, ""), nil)
	if err != nil {
		return nil, err
	}
	body, err := proxy.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	resources := []CdnHttpResource{}
	err = json.Unmarshal(body, &resources)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (proxy *ConfigurationApiProxy) CreateHttpResource(httpResource CdnHttpResource) (*CdnHttpResourceCreated, error) {
	rb, err := json.Marshal(httpResource)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(ConfigurationApiURL, proxy.AccountName, ""), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := proxy.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	response := CdnHttpResourceCreated{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	// TODO: make custom error
	if response.Status != "accept" {
		return nil, fmt.Errorf("message: %s, description: %s", response.Message, response.Description)
	}
	return &response, nil
}

func (proxy *ConfigurationApiProxy) GetHttpResource(resource_id string) (CdnHttpResource, error) {
	resource := CdnHttpResource{}
	req, err := http.NewRequest("GET", fmt.Sprintf(ConfigurationApiURL, proxy.AccountName, resource_id), nil)
	if err != nil {
		return resource, err
	}
	body, err := proxy.MakeRequest(req)
	if err != nil {
		return resource, err
	}

	err = json.Unmarshal(body, &resource)
	if err != nil {
		return resource, err
	}
	return resource, nil
}

// TODO: change response struct (without resource id)
func (proxy *ConfigurationApiProxy) UpdateHttpResource(httpResource CdnHttpResource, resource_id string) (*CdnHttpResourceCreated, error) {
	rb, err := json.Marshal(httpResource)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(ConfigurationApiURL, proxy.AccountName, resource_id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, err := proxy.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	response := CdnHttpResourceCreated{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "accept" {
		return nil, fmt.Errorf("message: %s, description: %s", response.Message, response.Description)
	}
	return &response, nil

}
func (proxy *ConfigurationApiProxy) DeactivateHttpResource(resource_id string) error {
	active := false
	httpResource := CdnHttpResource{Active: &active}

	rb, err := json.Marshal(httpResource)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf(ConfigurationApiURL, proxy.AccountName, resource_id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	body, err := proxy.MakeRequest(req)
	if err != nil {
		return err
	}

	response := CdnHttpResourceCreated{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Status != "accept" {
		return fmt.Errorf("message: %s, description: %s", response.Message, response.Description)
	}
	return nil
}
