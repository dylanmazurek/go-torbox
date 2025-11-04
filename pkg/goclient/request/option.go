package request

import "net/url"

type options struct {
	baseUrl string

	headers   map[string]string
	urlParams *url.Values
	body      any
	bodyType  BodyType
}

func defaultOptions() options {
	defaultOptions := options{}

	return defaultOptions
}

type Option func(*options)

func WithBaseURL(u string) Option {
	return func(o *options) {
		o.baseUrl = u
	}
}

func WithURLParams(p *url.Values) Option {
	return func(o *options) {
		o.urlParams = p
	}
}

func WithHeaders(h map[string]string) Option {
	return func(o *options) {
		o.headers = h
	}
}

type BodyType string

const (
	BodyTypeJSON BodyType = "json"
	BodyTypeForm BodyType = "form"
	BodyTypeFile BodyType = "file"
)

func WithJSONBody(b any) Option {
	return func(o *options) {
		o.body = b
		o.bodyType = BodyTypeJSON
	}
}

func WithFormBody(b any) Option {
	return func(o *options) {
		o.body = b
		o.bodyType = BodyTypeForm
	}
}

func WithFileBody(b any) Option {
	return func(o *options) {
		o.body = b
		o.bodyType = BodyTypeFile
	}
}
