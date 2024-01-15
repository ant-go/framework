package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/ant-go/framework/contracts/icodec"
	"github.com/ant-go/framework/contracts/ierror"
	"github.com/ant-go/framework/contracts/ihttp"
)

type Response struct {
	httpResponse *http.Response
	bodyBuffer   *bytes.Buffer
}

func NewResponse(httpResponse *http.Response) (response *Response, err error) {
	originBody := httpResponse.Body
	defer func() { _ = originBody.Close() }()

	bodyBuffer := new(bytes.Buffer)
	if _, err = bodyBuffer.ReadFrom(originBody); err != nil {
		err = ierror.New(ihttp.ErrReadBodyFailed, "err", err)
		return
	}

	httpResponse.Body = io.NopCloser(bodyBuffer)

	response = &Response{httpResponse: httpResponse, bodyBuffer: bodyBuffer}

	return
}

func (r Response) StatusCode() int {
	return r.httpResponse.StatusCode
}

func (r Response) StatusText() string {
	return r.httpResponse.Status
}

func (r Response) OK() bool {
	return r.httpResponse.StatusCode >= 200 && r.httpResponse.StatusCode < 300
}

func (r Response) Forbidden() bool {
	return r.httpResponse.StatusCode == http.StatusForbidden
}

func (r Response) NotFound() bool {
	return r.httpResponse.StatusCode == http.StatusNotFound
}

func (r Response) ClientError() bool {
	return r.httpResponse.StatusCode >= 400 && r.httpResponse.StatusCode < 500
}

func (r Response) ServerError() bool {
	return r.httpResponse.StatusCode >= 500 && r.httpResponse.StatusCode < 600
}

func (r Response) Redirect() bool {
	switch r.httpResponse.StatusCode {
	case 201, 301, 302, 303, 307, 308:
		return true
	}
	return false
}

func (r Response) Unmarshal(codecInterface icodec.CodecInterface, target any) (err error) {
	return codecInterface.Unmarshal(r.bodyBuffer.Bytes(), target)
}

func (r Response) HttpResponse() *http.Response {
	return r.httpResponse
}

func (r Response) ResponseHeader() http.Header {
	return r.httpResponse.Header
}

func (r Response) ResponseBody() *bytes.Buffer {
	return r.bodyBuffer
}
