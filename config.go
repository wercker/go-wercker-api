package wercker

import (
	"net/http"

	"github.com/wercker/go-wercker-api/credentials"
)

var defaultCredentialsProvider = credentials.NewMultiProvider(
	&credentials.EnvProvider{},
	credentials.Anonymous(),
)

var defaultConfig = &Config{
	Creds:      defaultCredentialsProvider,
	Endpoint:   "https://app.wercker.com",
	HTTPClient: http.DefaultClient,
}

// Config contains all configurable settings which will be used when making
// requests
type Config struct {
	Creds      credentials.Provider
	Endpoint   string
	HTTPClient *http.Client
}

// Copy will create a shallow copy of the Copy object
func (o *Config) Copy() *Config {
	newConfig := &Config{
		Creds:      o.Creds,
		Endpoint:   o.Endpoint,
		HTTPClient: o.HTTPClient,
	}
	return newConfig
}

// Merge creates a new shallow copy of o, and copies all non empty values from
// config to the copy. If config is nil, than o will be returned.
func (o *Config) Merge(config *Config) *Config {
	if config == nil {
		return o
	}

	newConfig := o.Copy()

	if config.Creds != nil {
		newConfig.Creds = config.Creds
	}

	if config.Endpoint != "" {
		newConfig.Endpoint = config.Endpoint
	}

	if config.HTTPClient != nil {
		newConfig.HTTPClient = config.HTTPClient
	}

	return newConfig
}
