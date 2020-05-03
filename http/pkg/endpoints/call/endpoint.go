package call

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"

	"github.com/open-integration/core/pkg/logger"
)

type (
	CallOptions struct {
		Context   context.Context
		LoggerFD  string
		Arguments *CallArguments
	}
)

func Call(opt CallOptions) (*CallReturns, error) {
	log := logger.New(&logger.Options{
		FilePath: opt.LoggerFD,
	})

	u, err := urlpkg.Parse(opt.Arguments.URL)
	if err != nil {
		return nil, err
	}

	var body io.ReadCloser
	if opt.Arguments.Content != nil {
		body = ioutil.NopCloser(bytes.NewReader([]byte(*opt.Arguments.Content)))
	}

	client := http.Client{}
	headers := http.Header{}
	for _, h := range opt.Arguments.Headers {
		headers.Set(*h.Name, *h.Value)
	}

	req := &http.Request{
		Method:     opt.Arguments.Verb,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     headers,
		Body:       body,
		Host:       u.Host,
	}
	req.Method = opt.Arguments.Verb

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	log.Debug("Request returns", "status", resp.Status)
	respheaders := []Header{}
	for name, value := range resp.Header {
		respheaders = append(respheaders, Header{
			Name:  &name,
			Value: &value[0],
		})
	}
	defer resp.Body.Close()
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &CallReturns{
		Status:  resp.StatusCode,
		Headers: respheaders,
		Body:    string(bodyData),
	}, nil
}
