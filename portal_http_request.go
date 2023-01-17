package libportal

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

type ConstantError string

func (ce ConstantError) Error() string {
	return string(ce)
}

func (ce ConstantError) Wrap(err error) error {
	if err == nil {
		return nil
	}
	return errors.WithMessage(err, ce.Error())
}

func (ce ConstantError) WithMessage(msg string) error {
	return errors.WithMessage(ce, msg)
}

const (
	ErrMethodNotSupported      ConstantError = "method not supported"
	ErrContentTypeNotSupported ConstantError = "content-type not supported"
	ErrInvalidURL              ConstantError = "invalid url"
	ErrEmpty                   ConstantError = "empty"
	ErrInvalidMediaType        ConstantError = "invalid media type"
)

const (
	applicationJSON           = "application/json"
	applicationFormURLEncoded = "application/x-www-form-urlencoded"
	multipartFormData         = "multipart/form-data"
)

type portalHTTPBuilder func(portalHTTPRequest *PortalHTTPRequest) (*http.Request, error)

var methodBuilderMap = map[string]portalHTTPBuilder{
	http.MethodGet:  buildGet,
	http.MethodPost: buildPost,
}

type encoderBuilder func(params map[string]string) (io.Reader, string, error)

var contentTypeMap = map[string]encoderBuilder{
	multipartFormData:         buildFormData,
	applicationJSON:           buildJSON,
	applicationFormURLEncoded: buildFormUrlEncoded,
}

func getHTTPBuilder(method string) (portalHTTPBuilder, error) {
	builder, ok := methodBuilderMap[method]
	if !ok {
		return nil, ErrMethodNotSupported
	}
	return builder, nil
}

func getEncoderBuilder(contentType string) (encoderBuilder, error) {
	builder, ok := contentTypeMap[contentType]
	if !ok {
		return nil, ErrContentTypeNotSupported
	}
	return builder, nil
}

type PortalHTTPRequest struct {
	Method      string            `json:"method" bson:"method"`
	Url         string            `json:"url" bson:"url"`
	ContentType string            `json:"content_type" bson:"content_type"`
	Headers     map[string]string `json:"headers" bson:"headers"`
	Params      map[string]string `json:"params" bson:"params"`
}

func (p *PortalHTTPRequest) Build() (*http.Request, error) {
	if _, err := url.Parse(p.Url); err != nil {
		return nil, ErrInvalidURL.Wrap(err)
	}
	builder, err := getHTTPBuilder(p.Method)
	if err != nil {
		return nil, err
	}
	return builder(p)
}

func buildGet(r *PortalHTTPRequest) (*http.Request, error) {
	return http.NewRequest(r.Method, r.Url, nil)
}
func buildPost(r *PortalHTTPRequest) (*http.Request, error) {
	if r.ContentType == "" {
		return nil, ErrEmpty.WithMessage("content-type")
	}

	ct, _, err := mime.ParseMediaType(r.ContentType)
	if err != nil {
		return nil, ErrInvalidMediaType.Wrap(err)
	}

	var body io.Reader = nil
	if r.Params != nil {
		builder, err := getEncoderBuilder(ct)
		if err != nil {
			return nil, ErrInvalidMediaType.Wrap(err)
		}
		body, ct, err = builder(r.Params)
		if err != nil {
			return nil, ErrInvalidMediaType.Wrap(err)
		}
	}

	req, err := http.NewRequest(r.Method, r.Url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", ct)
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	return req, nil
}

func buildJSON(params map[string]string) (io.Reader, string, error) {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(params); err != nil {
		return nil, "", err
	}
	return &b, applicationJSON, nil
}
func buildFormData(params map[string]string) (io.Reader, string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	for fieldName, value := range params {
		if err := writer.WriteField(fieldName, value); err != nil {
			return nil, "", err
		}
	}
	return &b, writer.FormDataContentType(), writer.Close()
}
func buildFormUrlEncoded(params map[string]string) (io.Reader, string, error) {
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}
	return strings.NewReader(data.Encode()), applicationFormURLEncoded, nil
}
