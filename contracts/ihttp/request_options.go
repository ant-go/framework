package ihttp

import (
	"net/http"
	"net/url"

	"github.com/ant-go/framework/contracts/icodec"
)

type BodyType int

const (
	BodyTypeRaw BodyType = iota
	BodyTypeForm
	BodyTypeJson
)

const (
	HttpRequestErrorNewRequestFailed = "http: new request failed"
	HttpRequestErrorDoFailed         = "http: do failed"
)

type RequestOption func(option *RequestOptions)

type RequestOptions struct {
	Query        url.Values
	Codec        icodec.CodecInterface
	Body         any
	BodyType     BodyType
	RawBody      []byte
	MaxRetries   int
	RetryChecker func(currentRetryTime int, currentErr error, response *http.Response) (err error)
}

func WithCodec(c icodec.CodecInterface) RequestOption {
	return func(option *RequestOptions) {
		if c != nil {
			option.Codec = c
		}
	}
}

func WithParams(params url.Values) RequestOption {
	return func(option *RequestOptions) {
		option.Query = params
	}
}

func WithFormBody(body any) RequestOption {
	return func(option *RequestOptions) {
		option.Body = body
		option.BodyType = BodyTypeForm
	}
}

func WithJsonBody(body any) RequestOption {
	return func(option *RequestOptions) {
		option.Body = body
		option.BodyType = BodyTypeJson
	}
}

func WithRawBody(rawBody []byte) RequestOption {
	return func(option *RequestOptions) {
		if rawBody != nil {
			option.RawBody = rawBody
		}
	}
}
