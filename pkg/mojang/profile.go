package mojang

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

const (
	profilePropertyTextures = "textures"
)

var (
	ErrResponseParse  = errors.New("failed to parse response data")
	ErrPropertyDecode = errors.New("failed to decode profile property")
)

type Profile struct {
	ID         uuid.UUID         `json:"id"`
	Name       string            `json:"name"`
	Properties []ProfileProperty `json:"properties"`
}

func (api *API) GetProfile(uuid uuid.UUID) (*Profile, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	normalizedUUID := strings.Replace(uuid.String(), "-", "", -1)
	req.SetRequestURI(fmt.Sprintf("%s/session/minecraft/profile/%s", api.sessionServer, normalizedUUID))

	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.Set("Accept", "application/json")

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := api.client.Do(req, res); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrHTTPSend, err)
	}

	profile := &Profile{}
	if err := json.Unmarshal(res.Body(), profile); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrResponseParse, err)
	}

	return profile, nil
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
