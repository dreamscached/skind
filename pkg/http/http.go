package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrRequestFailed = errors.New("failed to send HTTP request")
)

type Client interface {
	Request(ctx context.Context, config Request) (Response, error)
	RequestJSON(ctx context.Context, method, url string, body []byte) (Response, error)
	Get(ctx context.Context, url string) (Response, error)
}

type Request struct {
	Method         string
	URL            string
	RequestHeaders map[string]string
	RequestBody    []byte
}

type Response struct {
	Status int
	Body   []byte
}

type netHttpClient struct {
	httpClient *http.Client
}

func NewDefaultClient() Client {
	return &netHttpClient{httpClient: &http.Client{}}
}

func (client *netHttpClient) Request(ctx context.Context, config Request) (Response, error) {
	req, err := http.NewRequestWithContext(ctx, config.Method, config.URL, bytes.NewReader(config.RequestBody))
	if err != nil {
		return Response{}, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	for key, value := range config.RequestHeaders {
		req.Header.Set(key, value)
	}

	res, err := client.httpClient.Do(req)
	if err != nil {
		return Response{}, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	return Response{
		Status: res.StatusCode,
		Body:   body,
	}, nil
}

func (client *netHttpClient) RequestJSON(ctx context.Context, method, url string, body []byte) (Response, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	if body != nil {
		headers["Content-Type"] = "application/json"
	}

	return client.Request(ctx, Request{
		Method:         method,
		URL:            url,
		RequestHeaders: headers,
		RequestBody:    body,
	})
}

func (client *netHttpClient) Get(ctx context.Context, url string) (Response, error) {
	return client.Request(ctx, Request{
		Method: http.MethodGet,
		URL:    url,
	})
}
