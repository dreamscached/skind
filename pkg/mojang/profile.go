package mojang

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/dreamscached/skind/pkg/http"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const (
	profilePropertyTextures = "textures"
)

var (
	ErrResponseParseFailed = errors.New("failed to parse response data")
	ErrPropertyDecode      = errors.New("failed to decode profile property")
)

type Profile struct {
	ID         uuid.UUID         `json:"id"`
	Name       string            `json:"name"`
	Properties []ProfileProperty `json:"properties"`
}

func (api *API) GetProfile(uuid uuid.UUID) (*Profile, error) {
	properUUID := strings.Replace(uuid.String(), "-", "", -1)
	apiEndpoint := fmt.Sprintf("%s/session/minecraft/profile/%s", api.sessionServer, properUUID)

	res, err := api.client.SendHTTP(http.Request{
		Method:      fasthttp.MethodGet,
		URL:         apiEndpoint,
		RequestBody: nil,
	})
	if err != nil {
		return nil, err
	}

	switch res.Status {
	case fasthttp.StatusOK:
		profile := &Profile{}
		if err = json.Unmarshal(res.Body, profile); err != nil {
			return nil, fmt.Errorf("%w: %s", ErrResponseParseFailed, err)
		}
		return profile, nil
	case fasthttp.StatusNoContent, fasthttp.StatusNotFound:
		return nil, fmt.Errorf("%s: %w", uuid.String(), ErrNotFound)
	case fasthttp.StatusBadRequest:
		return nil, ErrBadRequest
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnexpectedStatus, res.Status)
	}
}

type ProfileProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProfileTexturesData struct {
	Textures ProfileTexturesImages `json:"textures"`
}

type ProfileTexturesImages struct {
	Skin SkinTexture `json:"SKIN"`
	Cape SkinTexture `json:"CAPE"`
}

func (prop ProfileProperty) DecodeBase64JSON(ptr any) error {
	decoded, err := base64.StdEncoding.DecodeString(prop.Value)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrPropertyDecode, err)
	}

	if err = json.Unmarshal(decoded, ptr); err != nil {
		return fmt.Errorf("%w: %s", ErrPropertyDecode, err)
	}

	return nil
}

func (prop ProfileProperty) IsTextures() bool {
	return prop.Name == profilePropertyTextures
}

func (prop ProfileProperty) DecodeTextures() (ProfileTexturesData, error) {
	var textures ProfileTexturesData
	err := prop.DecodeBase64JSON(&textures)
	return textures, err
}
