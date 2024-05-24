package mojang

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

type UsernameUUID struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (api *API) GetUUID(username string) (*UsernameUUID, error) {
	apiEndpoint := fmt.Sprintf("%s/minecraft/profile/lookup/bulk/byname", api.minecraftServices)
	requestBody, _ := json.Marshal([]string{username})

	res, err := api.client.RequestJSON(context.Background(), http.MethodGet, apiEndpoint, requestBody)
	if err != nil {
		return nil, err
	}

	switch res.Status {
	case http.StatusOK:
		results := make([]*UsernameUUID, 0, 1)
		if err = json.Unmarshal(res.Body, &results); err != nil {
			return nil, fmt.Errorf("%w: %s", ErrResponseParseFailed, err)
		}
		if len(results) == 0 {
			return nil, fmt.Errorf("%s: %w", username, ErrNotFound)
		}
		return results[0], nil
	case http.StatusBadRequest:
		return nil, ErrBadRequest
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedStatus, res.Status)
	}
}
