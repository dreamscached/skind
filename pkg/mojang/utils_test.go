package mojang

import (
	"context"
	"time"

	"github.com/dreamscached/skind/pkg/http"
)

type MockHTTPClient struct {
	NextResponse http.Response
	NextError    error
	Delay        time.Duration
}

func (m *MockHTTPClient) RespondWithString(status int, response string) {
	m.NextResponse = http.Response{
		Status: status,
		Body:   []byte(response),
	}
	m.NextError = nil
}

func (m *MockHTTPClient) RespondWithError(err error) {
	m.NextResponse = http.Response{}
	m.NextError = err
}

func (m *MockHTTPClient) Request(ctx context.Context, _ http.Request) (http.Response, error) {
	if m.Delay > 0 {
		select {
		case <-time.After(m.Delay):
		case <-ctx.Done():
			return http.Response{}, ctx.Err()
		}
	}

	if m.NextError != nil {
		return http.Response{}, m.NextError
	}
	return m.NextResponse, nil
}

func (m *MockHTTPClient) RequestJSON(ctx context.Context, _, _ string, _ []byte) (http.Response, error) {
	return m.Request(ctx, http.Request{})
}

func (m *MockHTTPClient) Get(ctx context.Context, _ string) (http.Response, error) {
	return m.Request(ctx, http.Request{})
}
