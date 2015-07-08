package wercker

import "github.com/wercker/go-wercker-api/credentials"

var defaultCredentialsProvider = credentials.NewMultiProvider(
	&credentials.EnvProvider{},
	credentials.Anonymous(),
)

var defaultOptions = &Options{
	Creds:    defaultCredentialsProvider,
	Endpoint: "https://app.wercker.com",
	// HTTPClient:  http.DefaultClient,
}

type Options struct {
	Creds    credentials.Provider
	Endpoint string
}

func (o *Options) Copy() *Options {
	newOptions := &Options{
		Creds:    o.Creds,
		Endpoint: o.Endpoint,
	}
	return newOptions
}

func (o *Options) Merge(options *Options) *Options {
	if options == nil {
		return o
	}

	newOptions := o.Copy()

	if options.Creds != nil {
		newOptions.Creds = options.Creds
	}

	if options.Endpoint != "" {
		newOptions.Endpoint = options.Endpoint
	}

	return newOptions
}

// default options < client options < function options
