package mojang

import (
	"testing"
)

func TestProfileProperty_DecodeBase64JSON(t *testing.T) {
	t.Run("Test Wiki.vg example", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name:  "textures",
			Value: "ewogICJ0aW1lc3RhbXAiIDogMTcxNjQ4MjU4MTc1OCwKICAicHJvZmlsZUlkIiA6ICI0NTY2ZTY5ZmM5MDc0OGVlOGQ3MWQ3YmE1YWEwMGQyMCIsCiAgInByb2ZpbGVOYW1lIiA6ICJUaGlua29mZGVhdGgiLAogICJ0ZXh0dXJlcyIgOiB7CiAgICAiU0tJTiIgOiB7CiAgICAgICJ1cmwiIDogImh0dHA6Ly90ZXh0dXJlcy5taW5lY3JhZnQubmV0L3RleHR1cmUvNzRkMWUwOGIwYmI3ZTlmNTkwYWYyNzc1ODEyNWJiZWQxNzc4YWM2Y2VmNzI5YWVkZmNiOTYxM2U5OTExYWU3NSIKICAgIH0sCiAgICAiQ0FQRSIgOiB7CiAgICAgICJ1cmwiIDogImh0dHA6Ly90ZXh0dXJlcy5taW5lY3JhZnQubmV0L3RleHR1cmUvYjBjYzA4ODQwNzAwNDQ3MzIyZDk1M2EwMmI5NjVmMWQ2NWExM2E2MDNiZjY0YjE3YzgwM2MyMTQ0NmZlMTYzNSIKICAgIH0KICB9Cn0=",
		}

		if !testProperty.IsTextures() {
			t.Error("expected IsTextures() to be true")
		}

		textures, err := testProperty.DecodeTextures()
		if err != nil {
			t.Errorf("error decoding textures: %v", err)
		}

		if textures.Textures.Skin == (SkinTexture{}) {
			t.Error("expected skin SkinTexture to not be zero")
		}

		if textures.Textures.Skin.URL == "" {
			t.Error("expected skin SkinTexture to not be empty")
		}

		if textures.Textures.Cape == (SkinTexture{}) {
			t.Error("expected cape SkinTexture to not be zero")
		}

		if textures.Textures.Cape.URL == "" {
			t.Error("expected cape SkinTexture to not be empty")
		}
	})
}
