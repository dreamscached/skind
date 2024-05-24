package mojang

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestAPI_GetUUID(t *testing.T) {
	client := &MockHTTPClient{}
	api := MustNewAPI(WithHTTPClient(client))

	t.Run("Test OK response", func(t *testing.T) {
		testUUID, _ := uuid.Parse("069a79f444e94726a5befca90e38aaf5")
		testName := "Notch"

		client.RespondWithString(http.StatusOK, `[{
			"id": "069a79f444e94726a5befca90e38aaf5",
			"name": "Notch"
		}]`)

		res, err := api.GetUUID("Notch")
		if err != nil {
			t.Fatalf("error getting UUID: %s", err)
		}

		if res.Name != testName {
			t.Errorf("expected Name to be '%s', got '%s'", testName, res.Name)
		}

		if res.ID != testUUID {
			t.Errorf("expected ID to be '%s', got '%s'", testUUID, res.ID)
		}
	})

	t.Run("Test missing username", func(t *testing.T) {
		client.RespondWithString(http.StatusOK, `[]`)
		_, err := api.GetUUID("missingUsername")
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("expected ErrNotFound error, got %v", err)
		}
	})

	t.Run("Test bad username", func(t *testing.T) {
		client.RespondWithString(http.StatusBadRequest, "")
		_, err := api.GetUUID("")
		if !errors.Is(err, ErrBadRequest) {
			t.Errorf("expected ErrBadRequest error, got %v", err)
		}
	})

	t.Run("Test unexpected error", func(t *testing.T) {
		client.RespondWithString(http.StatusTeapot, "")
		_, err := api.GetUUID("")
		if !errors.Is(err, ErrUnexpectedStatus) {
			t.Errorf("expected ErrUnexpectedStatus error, got %v", err)
		}
	})
}
