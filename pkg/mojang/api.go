package mojang

import (
	"errors"
	"fmt"

	"github.com/dreamscached/skind/pkg/http"
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrUnexpectedStatus = errors.New("unexpected status code")
)

type API struct {
	sessionServer     string
	minecraftServices string
	client            http.Client
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
		client:            http.NewDefaultClient(),
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

func WithHTTPClient(client http.Client) APIOption {
	return func(api *API) error {
		api.client = client
		return nil
	}
}
