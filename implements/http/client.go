package http

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	"github.com/ant-go/framework/contracts"
	"github.com/ant-go/framework/contracts/ierror"
	"github.com/ant-go/framework/contracts/ihttp"
)

type Client struct {
	options *ihttp.ClientOptions
}

func NewClient(options ...ihttp.ClientOption) *Client {
	var o *ihttp.ClientOptions
	contracts.ApplyOptions(o, options)

	return &Client{options: o}
}

func (c Client) Do(ctx context.Context, method, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	var o *ihttp.RequestOptions
	contracts.ApplyOptions(o, options)

	u, err := url.Parse(c.options.BaseURI + uri)
	if err != nil {
		return
	}
	if len(o.Query) > 0 {
		if len(u.RawQuery) > 0 {
			u.RawQuery += "&"
		}
		u.RawQuery += o.Query.Encode()
	}
	uri = u.String()

	httpResponse, err := c.do(ctx, method, uri, o)
	if err != nil {
		return
	}

	return NewResponse(httpResponse)
}

func (c Client) do(ctx context.Context, method, uri string, o *ihttp.RequestOptions) (resp *http.Response, err error) {
	var contentType string
	switch o.BodyType {
	case ihttp.BodyTypeRaw:
		contentType = ihttp.ContentTypeEmpty
	case ihttp.BodyTypeForm:
		o.RawBody, err = o.Codec.Marshal(&o.Body)
		contentType = ihttp.ContentTypeForm
	case ihttp.BodyTypeJson:
		o.RawBody, err = o.Codec.Marshal(&o.Body)
		contentType = ihttp.ContentTypeJson
	}

	for i := 0; i < o.MaxRetries; i++ {
		request, e := http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(o.RawBody))
		if e != nil {
			if err = o.RetryChecker(i, ierror.New(ihttp.HttpRequestErrorNewRequestFailed, "error", e), nil); err != nil {
				return
			}
			continue
		}
		request.Header.Add(ihttp.HeaderContentType, contentType)

		response, e := c.options.HttpClient.Do(request)
		if e != nil {
			if err = o.RetryChecker(i, ierror.New(ihttp.HttpRequestErrorDoFailed, "error", e), response); err != nil {
				return
			}
			continue
		}

		resp = response
		break
	}
	return
}

func (c Client) Get(ctx context.Context, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	return c.Do(ctx, http.MethodGet, uri, options...)
}

func (c Client) Post(ctx context.Context, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	return c.Do(ctx, http.MethodPost, uri, options...)
}

func (c Client) Put(ctx context.Context, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	return c.Do(ctx, http.MethodPut, uri, options...)
}

func (c Client) Patch(ctx context.Context, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	return c.Do(ctx, http.MethodPatch, uri, options...)
}

func (c Client) Delete(ctx context.Context, uri string, options ...ihttp.RequestOption) (resp ihttp.Response, err error) {
	return c.Do(ctx, http.MethodDelete, uri, options...)
}
