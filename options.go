package wercker

import "github.com/wercker/go-wercker-api/credentials"

type Options struct {
	Creds    credentials.Provider
	Endpoint string
}

// default options < client options < function options
