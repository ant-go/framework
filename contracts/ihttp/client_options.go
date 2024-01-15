package ihttp

import (
	"net/http"
)

type ClientOption func(option *ClientOptions)

type ClientOptions struct {
	BaseURI    string
	HttpClient *http.Client
}

func WithBaseURI(baseURI string) ClientOption {
	return func(option *ClientOptions) {
		option.BaseURI = baseURI
	}
}

func WithHttpClient(cli *http.Client) ClientOption {
	return func(option *ClientOptions) {
		option.HttpClient = cli
	}
}
