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
	Request(config Request) (Response, error)
	RequestJSON(method, url string, body []byte) (Response, error)
	Get(url string) (Response, error)
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

func (f *fasthttpClient) Request(config Request) (Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(config.Method)
	req.SetRequestURI(config.URL)

	for key, value := range config.RequestHeaders {
		req.Header.Set(key, value)
	}

	if config.RequestBody != nil {
		req.SetBody(config.RequestBody)
	}

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	return f.handleRequest(req, res)
}

func (f *fasthttpClient) RequestJSON(method, url string, body []byte) (Response, error) {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"

	if body != nil {
		headers["Content-Type"] = "application/json"
	}

	return f.Request(Request{
		Method:         method,
		URL:            url,
		RequestHeaders: headers,
		RequestBody:    body,
	})
}

func (f *fasthttpClient) Get(url string) (Response, error) {
	return f.Request(Request{
		Method: fasthttp.MethodGet,
		URL:    url,
	})
}

func (f *fasthttpClient) handleRequest(req *fasthttp.Request, res *fasthttp.Response) (Response, error) {
	if err := fasthttp.Do(req, res); err != nil {
		return Response{}, fmt.Errorf("%w: %s", ErrRequestFailed, err)
	}
	return Response{Status: res.StatusCode(), Body: res.Body()}, nil
}
