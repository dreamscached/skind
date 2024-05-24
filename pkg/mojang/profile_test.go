package mojang

import (
	"testing"
)

func TestProfileProperty_IsTextures(t *testing.T) {
	t.Run("Test valid textures property", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name: "textures",
		}

		if !testProperty.IsTextures() {
			t.Error("expected IsTextures() to be true")
		}
	})

	t.Run("Test invalid textures property", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name: "notTextures",
		}

		if testProperty.IsTextures() {
			t.Error("expected IsTextures() to be false")
		}
	})
}

func TestProfileProperty_DecodeBase64JSON(t *testing.T) {
	t.Run("Test Wiki.vg example", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name:  "textures",
			Value: "ewogICJ0aW1lc3RhbXAiIDogMTcxNjQ4MjU4MTc1OCwKICAicHJvZmlsZUlkIiA6ICI0NTY2ZTY5ZmM5MDc0OGVlOGQ3MWQ3YmE1YWEwMGQyMCIsCiAgInByb2ZpbGVOYW1lIiA6ICJUaGlua29mZGVhdGgiLAogICJ0ZXh0dXJlcyIgOiB7CiAgICAiU0tJTiIgOiB7CiAgICAgICJ1cmwiIDogImh0dHA6Ly90ZXh0dXJlcy5taW5lY3JhZnQubmV0L3RleHR1cmUvNzRkMWUwOGIwYmI3ZTlmNTkwYWYyNzc1ODEyNWJiZWQxNzc4YWM2Y2VmNzI5YWVkZmNiOTYxM2U5OTExYWU3NSIKICAgIH0sCiAgICAiQ0FQRSIgOiB7CiAgICAgICJ1cmwiIDogImh0dHA6Ly90ZXh0dXJlcy5taW5lY3JhZnQubmV0L3RleHR1cmUvYjBjYzA4ODQwNzAwNDQ3MzIyZDk1M2EwMmI5NjVmMWQ2NWExM2E2MDNiZjY0YjE3YzgwM2MyMTQ0NmZlMTYzNSIKICAgIH0KICB9Cn0=",
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

	t.Run("Test real API example (skin, slim + cape)", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name:  "textures",
			Value: "ewogICJ0aW1lc3RhbXAiIDogMTcxNjUzNDQwMzcwNiwKICAicHJvZmlsZUlkIiA6ICI2ZmZmOTdmZWQzNWQ0MjVjOWZiOWMxMzU1YThmYjExNyIsCiAgInByb2ZpbGVOYW1lIiA6ICJkcmVhbXNjYWNoZWQiLAogICJ0ZXh0dXJlcyIgOiB7CiAgICAiU0tJTiIgOiB7CiAgICAgICJ1cmwiIDogImh0dHA6Ly90ZXh0dXJlcy5taW5lY3JhZnQubmV0L3RleHR1cmUvZTExZWFmNDlmNDNlMzA3YWQwMDE1NzRjYmI3MWYxMThkZTMxNWRlZDNlMzJiMzc2OWFlZGIyMDZhMzliNTZmYyIsCiAgICAgICJtZXRhZGF0YSIgOiB7CiAgICAgICAgIm1vZGVsIiA6ICJzbGltIgogICAgICB9CiAgICB9LAogICAgIkNBUEUiIDogewogICAgICAidXJsIiA6ICJodHRwOi8vdGV4dHVyZXMubWluZWNyYWZ0Lm5ldC90ZXh0dXJlLzIzNDBjMGUwM2RkMjRhMTFiMTVhOGIzM2MyYTdlOWUzMmFiYjIwNTFiMjQ4MWQwYmE3ZGVmZDYzNWNhN2E5MzMiCiAgICB9CiAgfQp9",
		}

		textures, err := testProperty.DecodeTextures()
		if err != nil {
			t.Errorf("error decoding textures: %v", err)
		}

		if textures.Textures.Skin == (SkinTexture{}) {
			t.Error("expected skin SkinTexture to not be zero")
		}

		if !textures.Textures.Skin.Slim() {
			t.Error("expected skin SkinTexture to be slim")
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

	t.Run("Test real API example (skin, standard)", func(t *testing.T) {
		testProperty := ProfileProperty{
			Name:  "textures",
			Value: "ewogICJ0aW1lc3RhbXAiIDogMTcxNjUzNDk1ODk2MiwKICAicHJvZmlsZUlkIiA6ICIwNjlhNzlmNDQ0ZTk0NzI2YTViZWZjYTkwZTM4YWFmNSIsCiAgInByb2ZpbGVOYW1lIiA6ICJOb3RjaCIsCiAgInRleHR1cmVzIiA6IHsKICAgICJTS0lOIiA6IHsKICAgICAgInVybCIgOiAiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS8yOTIwMDlhNDkyNWI1OGYwMmM3N2RhZGMzZWNlZjA3ZWE0Yzc0NzJmNjRlMGZkYzMyY2U1NTIyNDg5MzYyNjgwIgogICAgfQogIH0KfQ==",
		}

		textures, err := testProperty.DecodeTextures()
		if err != nil {
			t.Errorf("error decoding textures: %v", err)
		}

		if textures.Textures.Skin == (SkinTexture{}) {
			t.Error("expected skin SkinTexture to not be zero")
		}

		if textures.Textures.Skin.Slim() {
			t.Error("expected skin SkinTexture to be not slim")
		}

		if textures.Textures.Skin.URL == "" {
			t.Error("expected skin SkinTexture to not be empty")
		}

		if textures.Textures.Cape != (SkinTexture{}) {
			t.Error("expected cape SkinTexture to not zero")
		}
	})
}
