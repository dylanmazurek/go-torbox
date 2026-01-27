package torbox

type options struct {
	apiKey  string
	baseUrl string
}

func defaultOptions() options {
	defaultOptions := options{}

	return defaultOptions
}

type Option func(*options)

func WithAPIKey(i string) Option {
	return func(o *options) {
		o.apiKey = i
	}
}

func WithBaseURL(u string) Option {
	return func(o *options) {
		o.baseUrl = u
	}
}
