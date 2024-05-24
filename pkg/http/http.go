package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	ErrRequestFailed = errors.New("failed to send HTTP request")
)

type Client interface {
	SendHTTP(Request) (Response, error)
}

type Request struct {
	Method      string
	URL         string
	RequestBody []byte
}

type Response struct {
	Status int
	Body   []byte
}

type fasthttpClient struct{ fasthttp.Client }

func NewDefaultClient() Client {
	return &fasthttpClient{Client: fasthttp.Client{
		ReadTimeout:                   1 * time.Second,
		WriteTimeout:                  1 * time.Second,
		MaxIdleConnDuration:           1 * time.Hour,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: 1 * time.Hour,
		}).Dial,
	}}
}

func (f *fasthttpClient) SendHTTP(config Request) (Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(config.Method)
	req.SetRequestURI(config.URL)

	req.Header.Set("Accept", "application/json")

	if config.RequestBody != nil {
		req.Header.SetContentType("application/json")
		req.SetBody(config.RequestBody)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		return Response{}, fmt.Errorf("%w: %s", ErrRequestFailed, err)
	}

	return Response{Status: res.StatusCode(), Body: res.Body()}, nil
}
