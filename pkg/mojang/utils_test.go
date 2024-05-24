package mojang

import (
	"github.com/dreamscached/skind/pkg/http"
)

type MockHTTPClient struct {
	NextResponse http.Response
}

func (m *MockHTTPClient) RespondWithString(status int, response string) {
	m.NextResponse = http.Response{
		Status: status,
		Body:   []byte(response),
	}
}

func (m *MockHTTPClient) Request(_ http.Request) (http.Response, error) {
	defer func() { m.NextResponse = http.Response{} }()
	return m.NextResponse, nil
}

func (m *MockHTTPClient) RequestJSON(_, _ string, _ []byte) (http.Response, error) {
	return m.Request(http.Request{})
}

func (m *MockHTTPClient) Get(_ string) (http.Response, error) {
	return m.Request(http.Request{})
}
