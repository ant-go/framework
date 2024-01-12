package ihttp

import (
	"net/http"

	"github.com/ant-go/framework/contracts/icodec"
)

type RequestOption func(option *RequestOptions)

type RequestOptions struct {
	HttpRequest *http.Request
	Codec       icodec.CodecInterface
	Body        any
	RawBody     []byte
}

func WithRequest(req *http.Request) RequestOption {
	return func(option *RequestOptions) {
		if req != nil {
			option.HttpRequest = req
		}
	}
}

func WithCodec(c icodec.CodecInterface) RequestOption {
	return func(option *RequestOptions) {
		if c != nil {
			option.Codec = c
		}
	}
}

func WithBody(body any) RequestOption {
	return func(option *RequestOptions) {
		option.Body = body
	}
}

func WithRawBody(rawBody []byte) RequestOption {
	return func(option *RequestOptions) {
		if rawBody != nil {
			option.RawBody = rawBody
		}
	}
}
