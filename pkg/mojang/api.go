package mojang

import (
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	ErrHTTPSend = errors.New("failed to send HTTP request")
)

type API struct {
	sessionServer     string
	minecraftServices string
	client            *fasthttp.Client
}

func MustNewAPI(options ...APIOption) *API {
	api, err := NewAPI(options...)

	if err != nil {
		panic(err)
	}

	return api
}

func NewAPI(options ...APIOption) (*API, error) {
	api := newDefaultAPI()

	for _, option := range options {
		if err := option(api); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return api, nil
}

func newDefaultAPI() *API {
	return &API{
		sessionServer:     "https://sessionserver.mojang.com",
		minecraftServices: "https://api.minecraftservices.com",
		client: &fasthttp.Client{
			ReadTimeout:                   1 * time.Second,
			WriteTimeout:                  1 * time.Second,
			MaxIdleConnDuration:           1 * time.Hour,
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: true,
			Dial: (&fasthttp.TCPDialer{
				Concurrency:      4096,
				DNSCacheDuration: 1 * time.Hour,
			}).Dial,
		},
	}
}

type APIOption func(*API) error

func WithSessionServer(baseURL string) APIOption {
	return func(api *API) error {
		api.sessionServer = baseURL
		return nil
	}
}

func WithMinecraftServices(baseURL string) APIOption {
	return func(api *API) error {
		api.minecraftServices = baseURL
		return nil
	}
}

func WithHTTPClient(client *fasthttp.Client) APIOption {
	return func(api *API) error {
		api.client = client
		return nil
	}
}
