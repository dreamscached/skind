package mojang

import (
	"errors"
	"fmt"

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
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(fmt.Sprintf("%s/minecraft/profile/lookup/bulk/byname", api.minecraftServices))

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	payload, _ := json.Marshal([]string{username})
	req.SetBody(payload)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := api.client.Do(req, res); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrHTTPSend, err)
	}

	results := make([]*UsernameUUID, 0, 1)
	if err := json.Unmarshal(res.Body(), &results); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponseParse, err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrUsernameNotFound, username)
	}

	return results[0], nil
}
