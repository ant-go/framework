package ihttp

import (
	"bytes"
	"context"
	"net/http"

	"github.com/ant-go/framework/contracts/icodec"
)

type Client interface {
	Do(ctx context.Context, method, url string, options ...RequestOption) (resp Response, err error)

	Get(ctx context.Context, url string, options ...RequestOption) (resp Response, err error)
	Post(ctx context.Context, url string, options ...RequestOption) (resp Response, err error)
	Put(ctx context.Context, url string, options ...RequestOption) (resp Response, err error)
	Patch(ctx context.Context, url string, options ...RequestOption) (resp Response, err error)
	Delete(ctx context.Context, url string, options ...RequestOption) (resp Response, err error)
}

type ResponseStatus interface {
	StatusCode() int
	StatusText() string
}

type ResponseOK interface {
	OK() bool          // [200,200]
	Forbidden() bool   // [403,403]
	NotFound() bool    // [404,404]
	ClientError() bool // [400,499)
	ServerError() bool // [500,599)
	Redirect() bool    // (201,301,302,303,307,308)
}

type ResponseCodec interface {
	Unmarshal(codecInterface icodec.CodecInterface, target any) (err error)
}

type ResponseInfo interface {
	HttpResponse() *http.Response
	ResponseHeader() http.Header
	ResponseBody() *bytes.Buffer
}

type Response interface {
	ResponseStatus
	ResponseOK
	ResponseCodec
	ResponseInfo
}
