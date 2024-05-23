package mojang

const (
	SkinModelSlim = "slim"
)

type SkinTexture struct {
	URL      string          `json:"url"`
	Metadata SkinTextureMeta `json:"metadata"`
}

type SkinTextureMeta struct {
	Model string `json:"model"`
}

func (skin SkinTexture) Slim() bool {
	return skin.Metadata.Model == SkinModelSlim
}
