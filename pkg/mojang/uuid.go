package mojang

import (
	"errors"
	"fmt"

	"github.com/dreamscached/skind/pkg/http"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

var (
	ErrUsernameNotFound = errors.New("username not found")
)

type UsernameUUID struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (api *API) GetUUID(username string) (*UsernameUUID, error) {
	apiEndpoint := fmt.Sprintf("%s/minecraft/profile/lookup/bulk/byname", api.minecraftServices)
	requestBody, _ := json.Marshal([]string{username})

	res, err := api.client.SendHTTP(http.Request{
		Method:      fasthttp.MethodPost,
		URL:         apiEndpoint,
		RequestBody: requestBody,
	})
	if err != nil {
		return nil, err
	}

	results := make([]*UsernameUUID, 0, 1)
	if err := json.Unmarshal(res.Body, &results); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponseParseFailed, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrUsernameNotFound, username)
	}

	return results[0], nil
}
